package main

import (
	"io/ioutil"
	"os"

	"git.de-liker.com/Qs-F/expandup/lib/base"
)

func main() {
	file, _ := ioutil.ReadFile(os.Args[1])
	ioutil.WriteFile(os.Args[1], []byte(expandup.MustUp(string(file)).Compose()), 0600)
}
