package expandup

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Resolve determines the corresponding ident and content
// Each directry has own priority. This func resolves the order and uniquely defines the ident.
// Although ident can be the map (so that never needs to care about dups, but no map cuz of efficiency.
func (ps *paths) Resolve() []*Def {
	for _, path := range *ps {
		fext(path)
	}

	return nil
}

type path string

func NewPath(pathname string) path {
	p, err := filepath.Abs(pathname)
	if err != nil {
		return path(filepath.Clean(pathname))
	}
	return path(p)
}

type paths []path

type file struct {
	Path    path
	Content []byte
}

type files []*file

var NoFiles = &files{}

// If any error occurs for os.Open or filepath.Abs, like files not found,
// This cannot be tested
func fext(path path) *files {
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

		*ret = append(*ret, &file{Path: NewPath(filepath.Join(string(p), fi.Name())), Content: b})
	}

	if len(*ret) == 0 {
		return NoFiles
	}
	return ret
}

func (f *files) Len() int {
	return len(*f)
}

func (f *files) Less(i, j int) bool {
	return string((*f)[i].Path) > string((*f)[j].Path)
}
