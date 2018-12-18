package day18

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"time"
)

type acre int

const (
	OpenGround acre = iota
	Trees
	Lumberyard
)

func Solve(lines []string, partOne bool) string {

	currentWorld := make(map[Coordinate]acre, 0)
	for y, l := range lines {
		for x, a := range l {
			c := Coordinate{x, y}
			if a == '.' {
				currentWorld[c] = OpenGround
			} else if a == '|' {
				currentWorld[c] = Trees
			} else if a == '#' {
				currentWorld[c] = Lumberyard
			}
		}
	}
	maxX := len(lines[0])
	maxY := len(lines)

	minutes := 0
	printWorld(currentWorld, maxX, maxY)
	oldValue := -1
	start := time.Now()
	values := make([]int, 0)
	for {
		nextWorld := make(map[Coordinate]acre, 0)
		for c, a := range currentWorld {
			_, trees, lumberyards := countAdjacent(c, currentWorld)
			if a == OpenGround {
				if trees >= 3 {
					nextWorld[c] = Trees
				} else {
					nextWorld[c] = a
				}

			} else if a == Trees {
				if lumberyards >= 3 {
					nextWorld[c] = Lumberyard
				} else {
					nextWorld[c] = a
				}
			} else if a == Lumberyard {
				if trees >= 1 && lumberyards >= 1 {
					nextWorld[c] = a
				} else {
					nextWorld[c] = OpenGround
				}
			}
		}
		currentWorld = nextWorld
		minutes++
		//printWorld(currentWorld, maxX, maxY)
		if minutes%10000 == 0 {
			t := time.Now()
			elapsed := t.Sub(start)
			printWorld(currentWorld, maxX, maxY)
			value := calculateResourceValue(currentWorld)
			fmt.Printf("Value %v after %v in %v\n", value, minutes, elapsed)
			fmt.Printf("Diff %v\n", value-oldValue)
			fmt.Printf("Progress%v\n", float64(minutes)/float64(1000000000))
			oldValue = value
			if minutes > 10000 {
				if value == values[0] {
					break
				}
			}
			values = append(values, value)
		}
		if partOne && minutes == 10 {
			break
		}
		//     10000
		if minutes == 1000000000 {
			break
		}
	}

	if partOne {
		resourceValue := calculateResourceValue(currentWorld)
		return fmt.Sprint(resourceValue)
	}
	correctIndex := ((1000000000 - minutes) / 10000) % len(values)
	return fmt.Sprint(values[correctIndex])
}

func calculateResourceValue(acres map[Coordinate]acre) int {
	trees := 0
	lumberyards := 0
	for _, a := range acres {
		if a == Trees {
			trees++
		} else if a == Lumberyard {
			lumberyards++
		}
	}

	return trees * lumberyards
}

func printWorld(acres map[Coordinate]acre, maxX int, maxY int) {
	str := ""
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			a := acres[Coordinate{X: x, Y: y}]
			if a == OpenGround {
				str += "."
			} else if a == Trees {
				str += "|"
			} else if a == Lumberyard {
				str += "#"
			}
		}

		str += "\n"
	}

	fmt.Println(str)
}

func countAdjacent(c Coordinate, acres map[Coordinate]acre) (int, int, int) {
	openGround := 0
	trees := 0
	lumberyards := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}

			a, f := acres[c.Plus(Coordinate{X: x, Y: y})]
			if !f {
				continue
			}

			if a == OpenGround {
				openGround++
			} else if a == Trees {
				trees++
			} else if a == Lumberyard {
				lumberyards++
			}
		}
	}

	return openGround, trees, lumberyards

}
