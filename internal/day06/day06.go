package day06

import (
	"fmt"
	. "github.com/Dr-Horv/Advent-of-Code-2018/internal/pkg"
	"math"
)

func Solve(lines []string, partOne bool) string {
	var coordinates []Coordinate

	for _, l := range lines {
		coordinates = append(coordinates, ParseCoordinate(l))
	}

	maxV := -1
	if partOne {
		explored := make(map[Coordinate]int)
		for i, c := range coordinates {
			r := expand(i, c, coordinates, explored)
			if r > maxV {
				maxV = r
			}
		}
	} else {
		var x0 = math.MaxInt64
		var y0 = math.MaxInt64
		var x1 = 0
		var y1 = 0

		for _, c := range coordinates {
			x0 = Min(c.X, x0)
			x1 = Max(c.X, x1)

			y0 = Min(c.X, y0)
			y1 = Max(c.Y, y1)
		}

		topLeft := Coordinate{X: x0, Y: y0}
		bottomRight := Coordinate{X: x1, Y: y1}
		explored := make(map[Coordinate]bool)

		var i = 0
		for y := topLeft.Y; y <= bottomRight.Y; y++ {
			for x := topLeft.X; x <= bottomRight.X; x++ {
				r := expandSafe(i, Coordinate{X: x, Y: y}, coordinates, explored)
				if r > maxV {
					maxV = r
				}
				i++
			}
		}
	}

	return fmt.Sprintf("%v", maxV)
}

func expandSafe(id int, c Coordinate, coordinates []Coordinate, explored map[Coordinate]bool) int {
	_, found := explored[c]
	if found {
		return -1
	}
	return expandSafeHelper(id, c, coordinates, explored)
}

func expandSafeHelper(id int, c Coordinate, coordinates []Coordinate, explored map[Coordinate]bool) int {
	explored[c] = true
	var totalDist = 0
	for _, curr := range coordinates {
		totalDist += ManhattanDistance(curr, c)
		if totalDist >= 10000 {
			explored[c] = false
			return -1
		}
	}

	if totalDist >= 10000 {
		explored[c] = false
		return -1
	}

	c1, c2, c3, c4 := GetNeighbours(c)

	var sum = 1
	for _, cn := range []Coordinate{c1, c2, c3, c4} {
		_, found := explored[cn]
		if !found {
			r := expandSafeHelper(id, cn, coordinates, explored)
			if r != -1 {
				sum += r
			}
		}
	}

	return sum
}

func expand(id int, c Coordinate, coordinates []Coordinate, explored map[Coordinate]int) int {
	return expandHelper(id, c, coordinates, explored, 0)
}

func expandHelper(id int, c Coordinate, coordinates []Coordinate, explored map[Coordinate]int, depth int) int {
	if depth > 100000 {
		return -1
	}

	var minV = math.MaxInt64
	var min int
	for ci, curr := range coordinates {
		dist := ManhattanDistance(curr, c)
		if dist <= minV {
			if dist == minV {
				min = -1
			} else {
				min = ci
			}

			minV = dist
		}
	}

	explored[c] = min

	if min != id {
		return 0
	}

	c1, c2, c3, c4 := GetNeighbours(c)

	sum := 1

	for _, c := range []Coordinate{c1, c2, c3, c4} {
		_, found := explored[c]
		if !found {
			r := expandHelper(id, c, coordinates, explored, depth+1)
			if r == -1 {
				return -1
			}
			sum += r
		}
	}

	return sum
}
