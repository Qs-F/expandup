package expandup

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Resolve determines the corresponding ident and content
// Each directry has own priority. This func resolves the order and uniquely defines the ident.
// Although ident can be the map (so that never needs to care about dups, but no map cuz of efficiency.
func (ps *Paths) Resolve() []*Def {
	for _, path := range *ps {
	}
}

type Path string

type Paths []Path

type file struct {
	Path    Path
	Content []byte
}

type files []*file

var NoFiles = &files{}

// If any error occurs for os.Open or filepath.Abs, like files not found,
func fext(path Path) *files {
	p, err := filepath.Abs(string(path))
	if err != nil {
		return NoFiles
	}

	stat, err := os.Stat(p)
	if err != nil {
		return NoFiles
	}

	if !stat.IsDir() { // if given file is not directory (regular file)
		f, err := os.Open(p)
		if err != nil {
			return NoFiles
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			return NoFiles
		}

		return &files{&file{Path: path, Content: b}}
	}

	// if directory
	f, err := os.Open(p)
	if err != nil {
		return NoFiles
	}

	fis, err := f.Readdir(-1)
	if err != nil {
		return NoFiles
	}

	ret := &files{}
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}

		f, err := os.Open(fi.Name())
		if err != nil {
			continue
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			continue
		}

		*ret = append(*ret, &file{Path: Path(filepath.Join(string(p), fi.Name())), Content: b})
	}

	if len(*ret) == 0 {
		return NoFiles
	}
	return ret
}

// // Path returns the absolute path.
// // filepath.Abs returns error, then
// func Path(path string) []path {
// 	paths := []path{}
//
// }
