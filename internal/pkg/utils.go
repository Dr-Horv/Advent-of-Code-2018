package pkg

import (
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func ManhattanDistance(c1 Coordinate, c2 Coordinate) int {
	return Abs(c1.X-c2.X) + Abs(c1.Y-c2.Y)

}

func GetNeighbours(c Coordinate) (Coordinate, Coordinate, Coordinate, Coordinate) {
	c1 := Coordinate{X: c.X + 1, Y: c.Y}
	c2 := Coordinate{X: c.X - 1, Y: c.Y}
	c3 := Coordinate{X: c.X, Y: c.Y + 1}
	c4 := Coordinate{X: c.X, Y: c.Y - 1}
	return c1, c2, c3, c4
}

func ParseCoordinate(s string) Coordinate {
	components := strings.Split(s, ",")
	x, _ := strconv.Atoi(components[0])
	y, _ := strconv.Atoi(strings.TrimSpace(components[1]))
	return Coordinate{X: x, Y: y}
}

func Compare(i1 int, i2 int, operator func(int, int) bool) int {
	if operator(i1, i2) {
		return i1
	} else {
		return i2
	}
}

func Min(i1 int, i2 int) int {
	return Compare(i1, i2, func(i1 int, i int) bool {
		return i1 < i2
	})
}

func Max(i1 int, i2 int) int {
	return Compare(i1, i2, func(i1 int, i int) bool {
		return i1 > i2
	})
}
