package day11

import (
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"testing"
)

func TestCalculateSquareValues(t *testing.T) {
	grid := make(map[Coordinate]int)

	i := 1
	for y := 1; y <= 3; y++ {
		for x := 1; x <= 3; x++ {
			grid[Coordinate{X: x, Y: y}] = i
			i++
		}
	}

	values := CalculateSquareValues(Coordinate{1, 1}, grid, 3)
	if values[0] != 1 {
		t.Error("Expected 1, got ", values[0])
	}

	if values[1] != 12 {
		t.Error("Expected 7, got ", values[1])
	}

	if values[2] != 45 {
		t.Error("Expected 45, got ", values[2])
	}
}
