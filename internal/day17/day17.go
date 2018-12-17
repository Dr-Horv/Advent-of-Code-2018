package day17

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"math"
	"regexp"
	"strings"
)

type square int

const (
	Clay square = iota
	Water
)

func Solve(lines []string, partOne bool) string {

	var xAndYRange = regexp.MustCompile(`(?m)x=(\d+), y=(\d+)..(\d+)`)
	var yAndXRange = regexp.MustCompile(`(?m)y=(\d+), x=(\d+)..(\d+)`)

	world := make(map[Coordinate]square, 0)
	for _, l := range lines {
		groupsX := xAndYRange.FindStringSubmatch(l)
		if groupsX != nil {
			addClay(world, StrConv(groupsX[1]), StrConv(groupsX[1]), StrConv(groupsX[2]), StrConv(groupsX[3]))
		} else {
			groupsY := yAndXRange.FindStringSubmatch(l)
			addClay(world, StrConv(groupsY[2]), StrConv(groupsY[3]), StrConv(groupsY[1]), StrConv(groupsY[1]))
		}

	}

	_, _, minY, maxY := calculateBounds(world)
	hasVisited := make(map[Coordinate]bool, 0)
	spreadWater(world, hasVisited, Coordinate{X: 500, Y: 1}, maxY)
	printWorld(world)
	if partOne {
		return fmt.Sprint(countWater(world, minY))
	} else {
		return fmt.Sprint(countStableWater(world, hasVisited, minY))
	}
}

func countStableWater(squares map[Coordinate]square, hasVisited map[Coordinate]bool, minY int) interface{} {
	sum := 0
	for k, v := range squares {
		if k.Y < minY {
			continue
		}
		if v == Water && !hasVisited[k] {
			sum++
		}
	}

	return sum
}

func countWater(squares map[Coordinate]square, minY int) int {
	sum := 0
	for k, v := range squares {
		if k.Y < minY {
			continue
		}
		if v == Water {
			sum++
		}
	}

	return sum
}

func isTerminatingOrCanSpreadDown(squares map[Coordinate]square, hasVisited map[Coordinate]bool, coordinate Coordinate, maxY int) (bool, bool) {
	v, hasVisitedBefore := hasVisited[coordinate]

	if hasVisitedBefore {
		return true, v
	}

	if maxY < coordinate.Y {
		hasVisited[coordinate] = true
		return true, true
	}

	s, f := squares[coordinate]
	if f {
		if s == Clay {
			hasVisited[coordinate] = false
			return true, false
		}
	}

	squares[coordinate] = Water
	canSpread := spreadWater(squares, hasVisited, coordinate.Down(), maxY)
	if canSpread {
		hasVisited[coordinate] = true
		return true, true
	}
	return false, false

}

func spreadWater(squares map[Coordinate]square, hasVisited map[Coordinate]bool, coordinate Coordinate, maxY int) bool {
	isEnd, value := isTerminatingOrCanSpreadDown(squares, hasVisited, coordinate, maxY)
	if isEnd {
		return value
	}

	canSpreadRight := spreadRight(squares, hasVisited, coordinate.Right(), maxY)
	canSpreadLeft := spreadLeft(squares, hasVisited, coordinate.Left(), maxY)

	if canSpreadRight && !canSpreadLeft {
		updateMemoryToTheLeft(squares, hasVisited, coordinate.Left())
	}

	if !canSpreadRight && canSpreadLeft {
		updateMemoryToTheRight(squares, hasVisited, coordinate.Left())
	}

	if canSpreadRight || canSpreadLeft {
		hasVisited[coordinate] = true
		return true
	}

	hasVisited[coordinate] = false
	return false
}

func spreadRight(squares map[Coordinate]square, hasVisited map[Coordinate]bool, coordinate Coordinate, maxY int) bool {
	isEnd, value := isTerminatingOrCanSpreadDown(squares, hasVisited, coordinate, maxY)
	if isEnd {
		return value
	}

	canSpreadRight := spreadRight(squares, hasVisited, coordinate.Right(), maxY)

	if canSpreadRight {
		hasVisited[coordinate] = true
		return true
	}

	hasVisited[coordinate] = false
	return false
}

func spreadLeft(squares map[Coordinate]square, hasVisited map[Coordinate]bool, coordinate Coordinate, maxY int) bool {
	isEnd, value := isTerminatingOrCanSpreadDown(squares, hasVisited, coordinate, maxY)
	if isEnd {
		return value
	}

	canSpreadLeft := spreadLeft(squares, hasVisited, coordinate.Left(), maxY)

	if canSpreadLeft {
		hasVisited[coordinate] = true
		return true
	}

	hasVisited[coordinate] = false
	return false
}

func updateMemoryToTheRight(squares map[Coordinate]square, hasVisited map[Coordinate]bool, coordinate Coordinate) {
	for {
		s := squares[coordinate]
		if s == Water {
			hasVisited[coordinate] = true
		} else {
			break
		}

		coordinate = coordinate.Right()
	}
}

func updateMemoryToTheLeft(squares map[Coordinate]square, hasVisited map[Coordinate]bool, coordinate Coordinate) {
	for {
		s := squares[coordinate]
		if s == Water {
			hasVisited[coordinate] = true
		} else {
			break
		}

		coordinate = coordinate.Left()
	}
}

func addClay(squares map[Coordinate]square, x1 int, x2 int, y1 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			squares[Coordinate{X: x, Y: y}] = Clay
		}
	}
}

func calculateBounds(squares map[Coordinate]square) (int, int, int, int) {
	minX := math.MaxInt64
	maxX := math.MinInt64
	minY := math.MaxInt64
	maxY := math.MinInt64
	for c := range squares {
		minX = Min(c.X, minX)
		maxX = Max(c.X, maxX)
		minY = Min(c.Y, minY)
		maxY = Max(c.Y, maxY)
	}

	return minX, maxX, minY, maxY
}

func printWorld(squares map[Coordinate]square) {
	minX, maxX, minY, maxY := calculateBounds(squares)
	var strB strings.Builder
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			s, found := squares[Coordinate{X: x, Y: y}]
			if !found {
				strB.WriteString(" ")
			} else {
				if s == Clay {
					strB.WriteString("#")
				} else if s == Water {
					strB.WriteString("~")
				}
			}
		}
		strB.WriteString("\n")
	}

	fmt.Println(strB.String())
}
