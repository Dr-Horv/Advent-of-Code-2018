package day22

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
)

type CavernType int

const (
	rocky CavernType = iota
	narrow
	wet
)

type equipment int

const (
	climbingGear equipment = iota
	torch
	neither
)

type State struct {
	pos       Coordinate
	gear      equipment
	timeSpent int
}

func Solve(lines []string, partOne bool) string {
	erosionLevels := make(map[Coordinate]int)
	cavern := make(map[Coordinate]CavernType)
	target := Coordinate{X: 9, Y: 758}
	//target := Coordinate{X: 10, Y: 10}
	depth := 8103
	//depth := 510

	extra := 0

	if !partOne {
		extra = 800
	}

	for y := 0; y <= (target.Y + extra); y++ {
		for x := 0; x <= (target.X + extra); x++ {
			c := Coordinate{X: x, Y: y}
			erosionLevel := calculateErosionLevel(c, depth, erosionLevels, target)
			erosionLevels[c] = erosionLevel
			ct := erosionLevel % 3
			switch ct {
			case 0:
				cavern[c] = rocky
			case 1:
				cavern[c] = wet
			case 2:
				cavern[c] = narrow
			}
		}
	}

	printCavern(cavern, target)
	risk := calculateRisk(cavern, target)

	if partOne {
		return fmt.Sprint(risk)
	}

	state := State{Coordinate{}, torch, 0}
	goal := State{target, torch, 0}

	path, _ := Dijkstra(state, goal, cavern)

	fmt.Println(path)
	return fmt.Sprintln(path[len(path)-1])
}

func calculateRisk(types map[Coordinate]CavernType, target Coordinate) interface{} {
	risk := 0
	for y := 0; y <= target.Y; y++ {
		for x := 0; x <= target.X; x++ {
			t := types[Coordinate{X: x, Y: y}]
			switch t {
			case rocky:
				risk += 0
			case narrow:
				risk += 2
			case wet:
				risk += 1
			}
		}
	}
	return risk
}

func printCavern(types map[Coordinate]CavernType, target Coordinate) {
	str := ""
	for y := 0; y <= target.Y; y++ {
		for x := 0; x <= target.X; x++ {
			t := types[Coordinate{X: x, Y: y}]
			switch t {
			case rocky:
				str += "."
			case narrow:
				str += "|"
			case wet:
				str += "="
			}
		}

		str += "\n"
	}
	fmt.Println(str)
}

func calculateErosionLevel(coordinate Coordinate, depth int, erosionLevels map[Coordinate]int, target Coordinate) int {
	geologicIndex := calculateGeologicIndex(coordinate, erosionLevels, target)
	return (geologicIndex + depth) % 20183

}

func calculateGeologicIndex(coordinate Coordinate, erosionLevels map[Coordinate]int, target Coordinate) int {
	if coordinate == (Coordinate{}) {
		return 0
	}

	if coordinate == target {
		return 0
	}

	if coordinate.Y == 0 {
		return coordinate.X * 16807
	}

	if coordinate.X == 0 {
		return coordinate.Y * 48271
	}

	return erosionLevels[coordinate.Left()] * erosionLevels[coordinate.Up()]

}
