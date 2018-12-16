package day01

import "testing"

func TestSolve(t *testing.T) {

	answer := Solve([]string{"+1", "+1", "-2"}, true)

	if answer != "0" {
		t.Error("Expected 0, got ", answer)
	}

	answer = Solve([]string{"+3", "+3", "+4", "-2", "-4"}, false)

	if answer != "10" {
		t.Error("Expected 10, got ", answer)
	}

	answer = Solve([]string{"+7", "+7", "-2", "-7", "-4"}, false)

	if answer != "14" {
		t.Error("Expected 14, got ", answer)
	}
}
