package expandup

import "testing"

// single test for constructor
func TestNewDef(t *testing.T) {
}

func TestEq(t *testing.T) {
	target1 := "hello world"
	target2 := "hello world"

	sum1 := sum(target1)
	sum2 := sum(target2)

	if !eq(sum1, sum2) {
		t.Error("Comp failed")
	}
}
