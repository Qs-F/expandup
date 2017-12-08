package expandup

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestSplitDoc(t *testing.T) {
	testSlice := []string{
		`hello`,
		`this is me.
you must write your code,
in Go style.
`,
		`Go is \n super lang.\n`,
	}
	for _, v := range testSlice {
		t.Log(spew.Sdump(splitDoc(v)))
	}
}

type TestMarkerBuf []struct {
	InputN int
	InputM *MarkerBuf
	Must   *MarkerBuf
}

var TestCleanupMarkerBuf = TestMarkerBuf{
	{
		InputN: -1,
		InputM: &MarkerBuf{
			Kind:    inside_expandup,
			Name:    "Test",
			Content: []string{"hello", "world"},
		},
		Must: &MarkerBuf{
			Kind:    inside_common,
			Name:    "",
			Content: []string{},
		},
	},

	{
		InputN: 0,
		InputM: &MarkerBuf{
			Kind:    inside_expandup,
			Name:    "Test",
			Content: []string{"hello", "world"},
		},
		Must: &MarkerBuf{
			Kind:    inside_common,
			Name:    "",
			Content: []string{"world"},
		},
	},

	{
		InputN: 1,
		InputM: &MarkerBuf{
			Kind:    inside_expandup,
			Name:    "Test",
			Content: []string{"hello", "world"},
		},
		Must: &MarkerBuf{
			Kind:    inside_common,
			Name:    "",
			Content: []string{},
		},
	},
}

func TestCleanup(t *testing.T) {
	for _, test := range TestCleanupMarkerBuf {
		test.InputM.cleanup(test.InputN)
		if assert.Equal(t, test.InputM, test.Must) {
			t.Log(spew.Sdump(test.InputM))
		} else {
			t.Log(spew.Sdump(test.InputM))
		}
	}
}

type TestTypeCommit []struct {
	InputN int
	InputM *MarkerBuf
	MustM  *MarkerBuf
	MustB  *Block
}

var TestVarCommit = TestTypeCommit{
	{
		InputN: -1,
		InputM: &MarkerBuf{
			Kind:    inside_expandup,
			Name:    "Test",
			Content: []string{"hello", "world"},
		},
		MustM: &MarkerBuf{
			Kind:    inside_common,
			Name:    "",
			Content: []string{},
		},
		MustB: &Block{
			Name:    "Test",
			Content: []string{"hello", "world"},
		},
	},

	{
		InputN: 0,
		InputM: &MarkerBuf{
			Kind:    inside_expandup,
			Name:    "Test",
			Content: []string{"hello", "world"},
		},
		MustM: &MarkerBuf{
			Kind:    inside_common,
			Name:    "",
			Content: []string{"world"},
		},
		MustB: &Block{
			Name:    "Test",
			Content: []string{"hello"},
		},
	},

	{
		InputN: 3,
		InputM: &MarkerBuf{
			Kind:    inside_expandup,
			Name:    "Test",
			Content: []string{"hello", "world"},
		},
		MustM: &MarkerBuf{
			Kind:    inside_common,
			Name:    "",
			Content: []string{},
		},
		MustB: &Block{
			Name:    "Test",
			Content: []string{"hello", "world"},
		},
	},

	{
		InputN: 2,
		InputM: &MarkerBuf{
			Kind:    inside_expandup,
			Name:    "Test",
			Content: []string{"hello", "world", "rebuild", "fm"},
		},
		MustM: &MarkerBuf{
			Kind:    inside_common,
			Name:    "",
			Content: []string{"fm"},
		},
		MustB: &Block{
			Name:    "Test",
			Content: []string{"hello", "world", "rebuild"},
		},
	},
}

func TestCommit(t *testing.T) {
	for _, test := range TestVarCommit {
		b := test.InputM.commit(test.InputN)
		if assert.Equal(t, b, test.MustB) {
			t.Log(spew.Sdump(test.InputM))
		} else {
			t.Log(spew.Sdump(test.InputM))
		}
	}
}

func TestMustUp1(t *testing.T) {
	t.Log(MustUp(`
<!DOCTYPE html>
<html>
	<head>
		<!-- EXPANDUP RIOT -->
	</head>
	<body>
	</body>
</html>`).Compose())
}

func TestMustUp2(t *testing.T) {
	t.Log(MustUp(`
<!DOCTYPE html>
<html>
	<head>
		<!-- EXPANDUP RIOT -->
				<script src="riot.min.js"></script>
		<!-- (EXPANDUP END) -->
	</head>
	<body>
	</body>
</html>`).Compose())
}

func TestMustUp3(t *testing.T) {
	t.Log(MustUp(`
<!DOCTYPE html>
<html>
	<head>
		<!-- EXPANDUP RIOT -->
		<script src="riot.min.js"></script>
		<h1>Hello</h1>
		<!-- (EXPANDUP END) -->
	</head>
	<body>
	</body>
</html>
`).Compose())
}

func TestMustUp4(t *testing.T) {
	t.Log(MustUp(`
<!DOCTYPE html>
<html>
	<head>
		<!-- EXPANDUP RIOT -->
		<script src="riot.min.js"></script>
		<h1>Hello</h1>
		<!-- (EXPANDUP END) -->
		<!-- EXPANDUP RIOT -->
	</head>
	<body>
	</body>
</html>
`).Compose())
}

func TestMustUp5(t *testing.T) {
	t.Log(MustUp(`
<!DOCTYPE html>
<html>
	<head>
		<!-- EXPANDUP RECURSIVE -->
	</head>
	<body>
	</body>
</html>
`).Compose())

}
