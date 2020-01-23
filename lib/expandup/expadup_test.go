package expandup

import (
	"reflect"
	"testing"
)

func TestDocument(t *testing.T) {
	type T struct {
		I string
		O *Doc
	}

	tests := []T{
		{
			I: `hello
world

world`,
			O: &Doc{
				s: []string{"hello", "world", "", "world"},
			},
		},
	}

	for _, test := range tests {
		if d := Document(test.I); !reflect.DeepEqual(d, test.O) { // the order is significant
			t.Errorf("expect %v but got %v\n", test.O, d)
		} else {
			t.Logf("got %v\n", d)
		}
	}
}
