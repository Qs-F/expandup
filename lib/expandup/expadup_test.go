package expandup

import "testing"

// single test for constructor
func TestNewDef(t *testing.T) {
}

func TestComp(t *testing.T) {
	target1 := "hello world"
	target2 := "hello world"

	sum1 := Sum(target1)
	sum2 := Sum(target2)

	if !Comp(sum1, sum2) {
		t.Error("Comp failed")
	}
}
