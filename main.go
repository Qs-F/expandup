package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gitlab.com/Qs-F/expandup/lib/expandup"

	"github.com/Sirupsen/logrus"
)

// flag
// l: list of files to be expanduped
// f: indicate one file
// w: rewrite file
// d: display diff
// c: indicate config directory
// debug: show all error
// diffcmd: indicate diff command

var (
	flagList    = flag.Bool("l", false, "list files to be expanduped")
	flagFile    = flag.String("f", "", "indicate only one file or directory")
	flagWrite   = flag.Bool("w", false, "write result to files")
	flagDiff    = flag.Bool("d", false, "show diff")
	flagConfig  = flag.String("c", "", "indicate config directory")
	flagDebug   = flag.Bool("debug", false, "turn on debug mode (show all errors)")
	flagDiffCmd = flag.String("diffcmd", "", "indicate diff command (default is embeded diff)")
)

type File struct {
	Name    string // file name
	Content string // expanduped file content
	Raw     []byte // raw content of file
}

type Files []*File

func main() {
	flag.Parse()

	// compose slice of file
	// fsbuf means fs buffer. Only for listing files
	var fsbuf Files
	if *flagFile != "" {
		// if -f flag setted, use it
		if stat, err := os.Lstat(*flagFile); err != nil {
			handleErr(err)
			return
		} else {
			if stat.IsDir() { // indicated direcotry
				fsbuf, err = getFiles(*flagFile)
				if err != nil {
					handleErr(err)
					return
				}
			} else { // indicated one file
				body, err := ioutil.ReadFile(*flagFile)
				if err != nil {
					handleErr(err)
					return
				}
				f := &File{
					Name: *flagFile,
					Raw:  body,
				}
				fsbuf = append(fsbuf, f)
			}
		}
	} else {
		// get list of current directory files
		var err error
		fsbuf, err = getFiles("./")
		if err != nil {
			handleErr(err)
			return
		}
	}

	// compose fs from fsbuf.
	// fs consists of files needed to expandup. Other files are dropped from slice.
	var fs Files
	for _, v := range fsbuf {
		d, t, err := expandup.Up(string(v.Raw))
		if err != nil {
			handleErr(err)
			continue
		}
		if !t {
			continue
		} else {
			fs = append(fs, &File{
				Name:    v.Name,
				Raw:     v.Raw,
				Content: d.Compose(),
			})
		}
	}

	// option cmd swicther
	switch {
	case *flagList:
		for _, v := range fs {
			fmt.Println(v.Name)
		}
	case *flagWrite:
		for _, v := range fs {
			err := ioutil.WriteFile(v.Name, []byte(v.Content), 0600)
			if err != nil {
				handleErr(err)
				continue
			}
		}
	default:
		for _, v := range fs {
			fmt.Println(v.Content)
		}
	}
}

func handleErr(err error) {
	if flag.Parsed() {
		if *flagDebug {
			logrus.Error(err.Error())
		}
	}
	return
}

func getFiles(path string) (Files, error) {
	var fsbuf Files
	if files, err := ioutil.ReadDir(path); err != nil {
		return nil, err
	} else {
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if body, err := ioutil.ReadFile(filepath.Join(path, f.Name())); err != nil {
				handleErr(err)
				continue
			} else {
				// add list to fsbuf
				fsbuf = append(fsbuf, &File{
					Name: filepath.Join(path, f.Name()),
					Raw:  body,
				})
			}
		}
	}
	return fsbuf, nil
}
