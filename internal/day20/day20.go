package day20

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"strings"
	"time"
)

type DoorStatus int

const (
	Open DoorStatus = iota
	Wall
	Unknown
)

type Room struct {
	Pos   Coordinate
	Up    DoorStatus
	Down  DoorStatus
	Left  DoorStatus
	Right DoorStatus
}

func Solve(lines []string, partOne bool) string {

	world := make(map[Coordinate]*Room, 0)
	regex := lines[0]
	pos := Coordinate{}
	parseWorld(pos, -1, world, regex[1:len(regex)-1])

	printWorld(world)

	maxDoors := -1
	checks := len(world)
	checked := 0
	start := time.Now()
	roomsAbove1000 := 0
	for k := range world {
		path, err := AStar(pos, k, world)
		if err != nil {
			continue
		}

		steps := len(path) - 1
		if steps > maxDoors {
			maxDoors = steps
		}

		if steps >= 1000 {
			roomsAbove1000++
		}

		if checked%10 == 0 {
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Printf("Progress %v in %v\n", float64(checked)/float64(checks), elapsed)
		}
		checked++
	}

	return fmt.Sprintf("Furthest %v, above 1000 %v", maxDoors, roomsAbove1000)

}

func printWorld(rooms map[Coordinate]*Room) {
	minX := 9999
	minY := 9999
	maxX := -1
	maxY := -1

	for k := range rooms {
		minX = Min(minX, k.X)
		minY = Min(minY, k.Y)
		maxX = Max(maxX, k.X)
		maxY = Max(maxY, k.Y)
	}

	var str strings.Builder
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			room, f := rooms[Coordinate{X: x, Y: y}]
			if f && room.Up == Open {
				str.WriteString("#-")
			} else {
				str.WriteString("##")
			}

			if maxX == x {
				str.WriteString("#")
			}
		}
		str.WriteString("\n")
		for x := minX; x <= maxX; x++ {
			room, f := rooms[Coordinate{X: x, Y: y}]
			suffix := "."
			if x == 0 && y == 0 {
				suffix = "x"
			}

			if f && room.Left == Open {
				str.WriteString("|" + suffix)
			} else {
				str.WriteString("#" + suffix)
			}

			if maxX == x {
				str.WriteString("#")
			}
		}
		str.WriteString("\n")

		for x := minX; x <= maxX; x++ {
			if maxY == y {
				room, f := rooms[Coordinate{X: x, Y: y}]
				if f && room.Down == Open {

					str.WriteString("#-")

				} else {
					str.WriteString("##")
				}
				if maxX == x {
					str.WriteString("#")
				}
			}

		}

	}

	fmt.Printf("%v,%v,%v,%v\n", minX, minY, maxX, maxY)
	fmt.Println(str.String())

}

func parseWorld(pos Coordinate, going Direction, rooms map[Coordinate]*Room, s string) string {
	r, f := rooms[pos]

	if !f {
		r = &Room{pos, Unknown, Unknown, Unknown, Unknown}
	}

	switch going {
	case UP:
		r.Down = Open
	case DOWN:
		r.Up = Open
	case LEFT:
		r.Right = Open
	case RIGHT:
		r.Left = Open
	}

	rooms[pos] = r

	if len(s) == 0 {
		return s
	}

	n := s[0]

	switch n {
	case 'E':
		r.Right = Open
		return parseWorld(pos.Right(), RIGHT, rooms, s[1:])
	case 'N':
		r.Up = Open
		return parseWorld(pos.Up(), UP, rooms, s[1:])
	case 'S':
		r.Down = Open
		return parseWorld(pos.Down(), DOWN, rooms, s[1:])
	case 'W':
		r.Left = Open
		return parseWorld(pos.Left(), LEFT, rooms, s[1:])
	case '(':
		for {
			s = parseWorld(pos, -1, rooms, s[1:])
			if s[0] == '|' {
				if s[1] == ')' {
					s = s[2:]
					break
				}
			} else if s[0] == ')' {
				s = s[1:]
				break
			}
		}
	case '|':
		return s
	case ')':
		return s
	}

	return parseWorld(pos, -1, rooms, s)

}
