// package expandup(lib/base) is helper package of expandup main pkg.
package expandup

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/Qs-F/walkup"
	homedir "github.com/mitchellh/go-homedir"
)

const (
	_EXPANDUP_START_PREFIX = "<!-- EXPANDUP "
	_EXPANDUP_START_SUFFIX = " -->"
	_EXPANDUP_END          = "<!-- (EXPANDUP END) -->"
)

const (
	inside_common = iota
	inside_expandup
)

func trimLeftSpace(s string) string {
	return strings.TrimLeftFunc(s, unicode.IsSpace)
}

func trimLeftSpaces(s []string) []string {
	r := make([]string, len(s))
	for i, line := range s {
		r[i] = trimLeftSpace(line)
	}
	return r[:]
}

func combine(arg []string) string {
	return strings.Join(arg, "")
}

func splitDoc(s string) []string {
	return strings.SplitAfter(s, "\n")
}

func isStartLine(line string) (bool, string) {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, _EXPANDUP_START_PREFIX) && strings.HasSuffix(line, _EXPANDUP_START_SUFFIX) {
		name := strings.TrimPrefix(line, _EXPANDUP_START_PREFIX)
		name = strings.TrimSuffix(name, _EXPANDUP_START_SUFFIX)
		name = strings.TrimSpace(name)
		return true, name
	}
	return false, ""
}

func isEndLine(line string) bool {
	line = strings.TrimSpace(line)
	if strings.EqualFold(line, _EXPANDUP_END) {
		return true
	}
	return false
}

// MarkerBuf is used for preproccess of generating each Blocks.
type MarkerBuf struct {
	Kind    int      // Kind means the specific lines are in EXPANDUP, or not.
	Name    string   // Name means the lines are inside of EXPANDUP COMMAND
	Content []string // Content is the handling of lines,
}

func (m *MarkerBuf) set(kind int, name string, firstContent string) {
	m.Kind = kind
	m.Name = name
	m.Content = append(m.Content, firstContent)
}

func (m *MarkerBuf) add(content string) {
	m.Content = append(m.Content, content)
}

func (m *MarkerBuf) cleanup(num int) {
	if num < 0 {
		num = len(m.Content)
	} else {
		num++
	}
	m.Kind = inside_common
	m.Name = ""
	m.Content = m.Content[num:]
}

func (m *MarkerBuf) equalKind(kind int) bool {
	if m.Kind == kind {
		return true
	}
	return false
}

// Block is for divided file content.
// Each Blocks handle each lines.
// Name is command name. Lines with no command, Name becomes empty string.
// Content is lines.
type Block struct {
	Name    string   // Name means Content(lines) are inside the EXPANDUP COMMAND. COMMAND equals Name.
	Content []string // Content is each lines.
}

func (b *Block) combine() string {
	return strings.Join(b.Content, "")
}

func (m *MarkerBuf) commit(num int) *Block {
	b := &Block{}
	if num < 0 {
		b = &Block{
			Name:    m.Name,
			Content: m.Content,
		}
	} else {
		if num-1 >= len(m.Content) {
			num = len(m.Content) - 1
		}
		b = &Block{
			Name:    m.Name,
			Content: m.Content[:num+1],
		}
	}
	m.cleanup(num)
	return b
}

// type Document is slice of pointer struct of Block.
type Document []*Block

// Compose generates string content of pointer struct of Document.
func (d *Document) Compose() string {
	s := ""
	for _, v := range *d {
		s += v.combine()
	}
	return s
}

func parse(slice []string) *Document {
	markbuf := &MarkerBuf{}
	document := &Document{}
	for _, line := range slice {
		if b, name := isStartLine(line); b { // start line
			if markbuf.Kind != inside_common { // inside expandup
				*document = append(*document, markbuf.commit(0))
				*document = append(*document, markbuf.commit(-1))
			} else {
				*document = append(*document, markbuf.commit(-1))
			}
			markbuf.set(inside_expandup, name, line)
			continue
		}
		if isEndLine(line) {
			if markbuf.Kind != inside_common {
				markbuf.add(line)
				*document = append(*document, markbuf.commit(-1))
				continue
			}
		}
		markbuf.add(line)
	}
	if markbuf.Kind != inside_common {
		*document = append(*document, markbuf.commit(0))
		*document = append(*document, markbuf.commit(-1))
	} else {
		*document = append(*document, markbuf.commit(-1))
	}
	return document
}

func getFile(filename string) ([]byte, error) {
	path := ""
	current, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	list := walkup.Walkup(current, ".expandup", 0)
	if len(list) > 0 {
		path = filepath.Join(list[0], filename)
	}
	b, err := ioutil.ReadFile(path)
	if err == nil {
		return b, nil
	}
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	b, err = ioutil.ReadFile(filepath.Join(home, "/.expandup/", filename))
	if err != nil {
		return nil, err
	}
	return b, nil
}

// func useFile(filename string) (*Document, bool, error) {
// 	b, err := getFile(filename)
// 	if err != nil {
// 		return nil, false, err
// 	}
// 	d, t, err := Up(string(b))
// 	if err != nil {
// 		return nil, t, err
// 	}
// 	return d, t, nil
// }

func replace(b *Block) (*Block, bool, error) { // if replaced, return true
	var status bool
	if b.Name == "" {
		return nil, false, errors.New("No command name")
	}
	file, err := getFile(b.Name)
	if err != nil {
		return nil, false, err
	}
	// file, t, err := useFile(b.Name)
	// if err != nil {
	// 	return nil, t, err
	// }
	strfile := _EXPANDUP_START_PREFIX + b.Name + _EXPANDUP_START_SUFFIX + "\n" + string(file) + _EXPANDUP_END + "\n"
	m1 := md5.Sum([]byte(combine(trimLeftSpaces(splitDoc(strfile)))))
	m2 := md5.Sum([]byte(combine(trimLeftSpaces(b.Content))))
	s := []string{}
	if !bytes.Equal(m1[:], m2[:]) {
		for _, v := range splitDoc(strfile) {
			s = append(s, v)
		}
		b.Content = s
		status = true
	}
	return b, status, nil
}

// Up returns pointer struct of Docuemnt, document changed status, and error.
// If it has any changes, middle return item becomes true. On other cases, the value becomes false.
func Up(s string) (*Document, bool, error) {
	result := &Document{}
	blocks := parse(splitDoc(s))
	var err error
	var status bool
	for _, v := range *blocks {
		if v.Name != "" {
			var sta bool
			v, sta, err = replace(v)
			if status == false && sta == true {
				status = true
			}
			if err != nil {
				return nil, false, err
			}
		}
		*result = append(*result, v)
	}
	if err != nil {
		return nil, false, err
	}
	return result, status, nil
}

// MustUp returns only pointer struct of Document, with no error.
// This must be used only for easy buildable main package. DONNOT use inside your 3rd party packages.
// Handling errors are important. For general usages, Up must be used.
func MustUp(s string) *Document {
	d, b, err := Up(s)
	fmt.Println("| ", b, " | ", err, " |")
	return d

}
