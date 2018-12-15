package day15

import (
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"testing"
)

func TestLessThanReadingOrder(t *testing.T) {

	answer := LessThanReadingOrder(Coordinate{5, 1}, Coordinate{7, 2})

	if !answer {
		t.Error("Expected true, got ", answer)
	}
}
