package expandup

// comment

import (
	"os"
	"path/filepath"
	"strings"
)

type path struct {
	Name     string
	FullPath string
}

const NoPaths = []path{}

func Path(path string) []path {
	p, err := filepath.Abs(path)
	if err != nil {
		return NoPaths
	}

	f, err := os.Open(p)
	if err != nil {
		return NoPaths
	}

	stat, err := f.Stat()
	if err != nil {
		return NoPaths
	}

	paths := []path{}

	if !stat.IsDir() {
		paths = []path{{filepath.Base(p), p}}
	} else {
		
}

type Config struct {
	OpenPrefix  string
	OpenSuffix  string
	ClosePrefix string
	CloseSuffix string

	IgnoreIndents bool

	BaseFiles []Path
}

type Doc struct {
	s []string
}

func filter(s []string) (ret []string) {
	for _, v := range s {
		if v == "" {
			continue
		}
		ret = append(ret, v)
	}

	return ret
}

func Document(raw string) *Doc {
	d := &Doc{}
	d.s = strings.Split(raw, "\n")

	return d
}

type Sub struct {
	Name    string
	Content string
	SHA1    []byte
}
