package day20

import "testing"

func TestSolve(t *testing.T) {

	answer := Solve([]string{"line1"}, false)

	if answer != "expected" {
		t.Error("Expected SOMETHING, got ", answer)
	}
}
