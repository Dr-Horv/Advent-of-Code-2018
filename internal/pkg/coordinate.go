package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Direction int

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

func (d Direction) String() string {
	switch d {
	case LEFT:
		return "<"
	case DOWN:
		return "v"
	case RIGHT:
		return ">"
	case UP:
		return "^"
	}

	panic("No representation")
}

func (d Direction) TurnRight() Direction {
	switch d {
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	}

	panic("Can't turn")
}

func (d Direction) TurnLeft() Direction {
	switch d {
	case DOWN:
		return RIGHT
	case RIGHT:
		return UP
	case UP:
		return LEFT
	case LEFT:
		return DOWN
	}

	panic("Can't turn")
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%v, %v)", c.X, c.Y)
}

func (c Coordinate) Left() Coordinate {
	return Coordinate{c.X - 1, c.Y}
}

func (c Coordinate) Right() Coordinate {
	return Coordinate{c.X + 1, c.Y}
}

func (c Coordinate) Up() Coordinate {
	return Coordinate{c.X, c.Y - 1}
}

func (c Coordinate) Down() Coordinate {
	return Coordinate{c.X, c.Y + 1}
}

func (c Coordinate) Move(direction Direction) Coordinate {
	switch direction {
	case DOWN:
		return c.Down()
	case UP:
		return c.Up()
	case RIGHT:
		return c.Right()
	case LEFT:
		return c.Left()
	}

	panic(fmt.Sprintf("Fail invalid direction %v", direction))
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
