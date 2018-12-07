package day06

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coordinate struct {
	X int
	Y int
}

func Solve(lines []string, partOne bool) string {
	var coordinates []coordinate

	for _, l := range lines {
		coordinates = append(coordinates, parseCoordinate(l))
	}

	maxV := -1
	if partOne {
		for i, c := range coordinates {
			r := expand(i, c, coordinates)
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
			if c.X < x0 {
				x0 = c.X
			}

			if c.X > x1 {
				x1 = c.X
			}

			if c.Y < y0 {
				y0 = c.Y
			}

			if c.Y > y1 {
				y1 = c.Y
			}
		}

		topLeft := coordinate{x0, y0}
		bottomRight := coordinate{x1, y1}
		explored := make(map[coordinate]bool)

		var i = 0
		for y := topLeft.Y; y <= bottomRight.Y; y++ {
			for x := topLeft.X; x <= bottomRight.X; x++ {
				r := expandSafe(i, coordinate{x, y}, coordinates, explored)
				if r > maxV {
					maxV = r
				}
				i++
			}
		}
	}

	return fmt.Sprintf("%v", maxV)
}

func expandSafe(id int, c coordinate, coordinates []coordinate, explored map[coordinate]bool) int {
	_, found := explored[c]
	if found {
		return -1
	}
	return expandSafeHelper(id, c, coordinates, explored)
}

func expandSafeHelper(id int, c coordinate, coordinates []coordinate, explored map[coordinate]bool) int {
	explored[c] = true
	var totalDist = 0
	for _, curr := range coordinates {
		totalDist += manhattanDistance(curr, c)
		if totalDist >= 10000 {
			explored[c] = false
			return -1
		}
	}

	if totalDist >= 10000 {
		explored[c] = false
		return -1
	}

	c1, c2, c3, c4 := getNeighbours(c)

	var sum = 1
	for _, cn := range []coordinate{c1, c2, c3, c4} {
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

func getNeighbours(c coordinate) (coordinate, coordinate, coordinate, coordinate) {
	c1 := coordinate{c.X + 1, c.Y}
	c2 := coordinate{c.X - 1, c.Y}
	c3 := coordinate{c.X, c.Y + 1}
	c4 := coordinate{c.X, c.Y - 1}
	return c1, c2, c3, c4
}

func expand(id int, c coordinate, coordinates []coordinate) int {
	return expandHelper(id, c, coordinates, make(map[coordinate]int), 0)
}

func expandHelper(id int, c coordinate, coordinates []coordinate, explored map[coordinate]int, depth int) int {
	if depth > 100000 {
		return -1
	}

	var minV = math.MaxInt64
	var min int
	var none = false
	for ci, curr := range coordinates {
		dist := manhattanDistance(curr, c)
		if dist <= minV {
			if dist == minV {
				none = true
			} else {
				none = false
			}

			minV = dist
			min = ci
		}
	}

	if none {
		explored[c] = -1
		return 0
	} else {
		explored[c] = min
	}

	if min != id {
		return 0
	}

	c1, c2, c3, c4 := getNeighbours(c)

	sum := 1

	for _, c := range []coordinate{c1, c2, c3, c4} {
		r := checkCoordinate(id, c, coordinates, explored, depth)
		if r == -1 {
			return -1
		}
		sum += r
	}

	return sum
}

func checkCoordinate(id int, c coordinate, coordinates []coordinate, explored map[coordinate]int, depth int) int {
	_, found := explored[c]
	if !found {
		return expandHelper(id, c, coordinates, explored, depth+1)
	}
	return 0
}

func manhattanDistance(c1 coordinate, c2 coordinate) int {
	return abs(c1.X-c2.X) + abs(c1.Y-c2.Y)

}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func parseCoordinate(s string) coordinate {
	components := strings.Split(s, ",")
	x, _ := strconv.Atoi(components[0])
	y, _ := strconv.Atoi(strings.TrimSpace(components[1]))
	return coordinate{x, y}
}
