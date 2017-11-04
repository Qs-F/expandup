package expandup

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseDoc(t *testing.T) {
	testSlice := []string{
		`hello`,
		`this is me.
you must write your code,
in Go style.
`,
		`Go is \n super lang.\n`,
	}
	for _, v := range testSlice {
		t.Log(spew.Sdump(parseDoc([]byte(v))))
	}
}
