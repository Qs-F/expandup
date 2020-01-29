package expandup

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestFextSingle(t *testing.T) {
	basepath := "./_testdata/fext/"

	type T struct {
		Input path
		Must  *files
	}

	tests := []T{
		{
			Input: NewPath(filepath.Join(basepath, "./testfile")),
			Must: &files{
				&file{
					Path:    NewPath(filepath.Join(basepath, "./testfile")),
					Content: []byte{},
				},
			},
		},
	}

	for _, test := range tests {
		if result := fext(test.Input); !reflect.DeepEqual(result, test.Must) {
			t.Errorf("expect %v but got %v\n", test.Must, result)
		} else {
			t.Logf("got %v\n", result)
		}
	}
}
