package day03

import "testing"

func TestSolve(t *testing.T) {

	answer := Solve([]string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}, false)

	if answer != "4" {
		t.Error("Expected 4, got ", answer)
	}
}
