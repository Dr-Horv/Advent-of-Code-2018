package day06

import (
	"fmt"
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


	var x0 int = 99999
	var y0 int = 99999
	var x1 int = 0
	var y1 int = 0


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
			y1  = c.Y
		}
	}

	topLeft := coordinate{x0, y0}
	bottomRight := coordinate{x1, y1}

	closeMap := make(map[coordinate]int)

	for x := topLeft.X; x <= bottomRight.X; x++ {
		for y := topLeft.Y; y <= bottomRight.Y; y++ {
			var minV = 99999
			var min int
			compare := coordinate{x,y}
			for ci, curr := range coordinates {
				dist := manhattanDistance(curr, compare)
				if dist < minV {
					minV = dist
					min = ci
				}
			}

			closeMap[compare] = min
		}
	}


	countMap := make(map[int]int)

	for _,v := range closeMap {
		count := countMap[v]
		count++
		countMap[v] = count
	}

	maxOwned := -1

	for _, v := range countMap {
		if v > maxOwned {
			maxOwned = v
		}
	}

	return fmt.Sprintf("%v", maxOwned)
}

func manhattanDistance(c1 coordinate, c2 coordinate) int  {
	return abs(c1.X - c2.X) + abs(c1.Y - c2.Y)

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
	y, _ := strconv.Atoi(components[1])
	return coordinate{x, y}
}