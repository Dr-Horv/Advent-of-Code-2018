package day25

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"strings"
)

type Point4D struct {
	X int
	Y int
	Z int
	T int
}

func (p1 Point4D) manhattanDistance(p2 Point4D) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y) + Abs(p1.Z-p2.Z) + Abs(p1.T-p2.T)
}

func Solve(lines []string, partOne bool) string {

	points := make([]Point4D, 0)

	for _, l := range lines {
		trimmed := strings.TrimSpace(l)
		parts := strings.Split(trimmed, ",")
		p := Point4D{StrConv(parts[0]), StrConv(parts[1]), StrConv(parts[2]), StrConv(parts[3])}
		points = append(points, p)
	}

	//fmt.Println(points)

	constellations := make([][]Point4D, 0)

	for _, p := range points {
		inConstallation := false
		//fmt.Printf("Checking %v\n", p)
		for ci, c := range constellations {
			for _, pic := range c {
				dist := pic.manhattanDistance(p)
				//fmt.Printf("Distance %v between %v and %v\n", dist, pic, p)
				if dist <= 3 {
					inConstallation = true
					//	fmt.Printf("Adding %v to %v\n", p, constellations[ci])
					constellations[ci] = append(constellations[ci], p)
					break
				}
			}
			if inConstallation {
				break
			}
		}

		if !inConstallation {
			//fmt.Printf("New constallation %v\n", p)
			constellations = append(constellations, []Point4D{p})
		}
	}

	//fmt.Println(constellations)

	for {
		stable := mergeConstellationIfPossible(constellations)
		if stable {
			break
		}
	}

	sum := 0
	for _, c := range constellations {
		if len(c) > 0 {
			sum++
		}
	}

	//fmt.Println(constellations)

	return fmt.Sprint(sum)
}

func mergeConstellationIfPossible(constallations [][]Point4D) bool {
	stable := true
	for ci1, c1 := range constallations {
		for ci2, c2 := range constallations {
			if ci1 == ci2 {
				continue
			}

			inSame := false
			for _, p1 := range c1 {
				for _, p2 := range c2 {
					if p1.manhattanDistance(p2) <= 3 {
						inSame = true
						break
					}
				}
				if inSame {
					break
				}
			}

			if inSame {
				constallations[ci1] = append(constallations[ci1], constallations[ci2]...)
				constallations[ci2] = []Point4D{}
				stable = false
			}
		}
	}

	return stable
}
