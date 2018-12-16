package day02

import "testing"

func TestSolve(t *testing.T) {

	answer := Solve([]string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab"},
		true)

	if answer != "12" {
		t.Error("Expected 12, got ", answer)
	}

	answer = Solve([]string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz"},
		false)

	if answer != "fgij" {
		t.Error("Expected fgij, got ", answer)
	}
}
