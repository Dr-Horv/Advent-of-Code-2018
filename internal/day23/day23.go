package day23

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"math"
	"regexp"
	"sort"
	"time"
)

type nanobot struct {
	x int
	y int
	z int
	r int
}

var ORIGO = nanobot{0,0,0,0}

func Solve(lines []string, partOne bool) string {

	var re = regexp.MustCompile(`(?m)pos=<(-*\d+),(-*\d+),(-*\d+)>, r=(-*\d+)`)
	bots := make([]nanobot, 0)

	for _, l := range lines {
		match := re.FindStringSubmatch(l)
		x := StrConv(match[1])
		y := StrConv(match[2])
		z := StrConv(match[3])
		r := StrConv(match[4])

		bots = append(bots, nanobot{x, y, z, r})
	}

	if partOne {
		inRange := botsInRangeOfStrongest(bots)
		return fmt.Sprint(inRange)
	}

	intersections := make(map[nanobot]int, 0)
	for _, b1 := range bots {
		for _, b2 := range bots {

			if hasIntersection(b1, b2) {
				curr1, found1 := intersections[b1]
				if !found1 {
					curr1 = 0
				}
				intersections[b1] = curr1 + 1
			}
		}
	}

	sort.Slice(bots, func(i, j int) bool {
		n1 := bots[i]
		n2 := bots[j]
		n1s := intersections[n1]
		n2s := intersections[n2]

		if n1s > n2s {
			return true
		} else if n1s == n2s {
			return manhattanDistance(ORIGO, n1) < manhattanDistance(ORIGO, n2)
		} else {
			return false
		}
	})


	fmt.Println(intersections)
	fmt.Println(len(intersections))
	fmt.Println(bots)
	fmt.Println(intersections[bots[0]])

	n := search1(bots[0], intersections[bots[0]], bots)

	fmt.Println(n)
	fmt.Println(manhattanDistance(n, ORIGO))

	return ""
}

func search1(curr nanobot, shouldBeInRange int, bots []nanobot) nanobot {
	fmt.Println("search1")
	minX := curr.x - curr.r
	minY := curr.y - curr.r
	minZ := curr.z - curr.r
	maxX := curr.x + curr.r
	maxY := curr.y + curr.r
	maxZ := curr.z + curr.r

	bestInRange := 0
	for x := minX; x <= maxX; x += 10 {
		for y := minY; y <= maxY; y += 10 {
			for z := minZ; z <= maxZ; z += 10 {
				test := nanobot{x,y,z, 0}
				inRange := 0
				for _, b := range bots {
					if isInRange(test, b) {
						inRange++
					}
				}
				if bestInRange < inRange {
					bestInRange = inRange
					curr = test
				}

			}
		}
	}

	fmt.Println("search1 2")

	for {
		neighbours := getNeighbours(curr)
		found := false
		for _, n := range neighbours {
			if manhattanDistance(n, ORIGO) >= manhattanDistance(curr, ORIGO) {
				continue
			}

			inRange := 0
			for _, b := range bots {
				if isInRange(n, b) {
					inRange++
				}
			}

			fmt.Println(inRange)
			fmt.Println(shouldBeInRange)
			if inRange == shouldBeInRange {
				curr = n
				fmt.Println("Found better")
				found = true
			}
		}

		if !found {
			return curr
		}
	}
}

func hasIntersection(n1 nanobot, n2 nanobot) bool {
	dx := n1.x - n2.x
	dy := n1.y - n2.y
	dz := n1.z - n2.z
	length := manhattanDistance(n1, n2)
	if isInRange(n1, n2) || isInRange(n2, n1) {
		return true
	} else if length > (n1.r) + (n2.r) {
		return false
	}



	l := math.Sqrt( float64(dx)*float64(dx) + float64(dy)*float64(dy) + float64(dz)*float64(dz) )
	dxf := float64(dx) / l
	dyf := float64(dy) / l
	dzf := float64(dz) / l


	for diff := -2.0; diff < 2.1; diff += 0.1 {
		currX := float64(n1.x) + (dxf)*float64(n1.r) + diff
		currY := float64(n1.y) + (dyf)*float64(n1.r) + diff
		currZ := float64(n1.z) + (dzf)*float64(n1.r) + diff
		fakeBot := nanobot{int(currX), int(currY), int(currZ), 0}
		if isInRange(fakeBot, n1) && isInRange(fakeBot, n2) {
			return true
		}
	}


	return false

}

func search(pos nanobot, bots []nanobot) nanobot {

	for {
		currentInRange := 0
		for _, b := range bots {
			if isInRange(pos, b) {
				currentInRange++
			}
		}

		neighbours := getNeighbours(pos)

		best := pos
		found := false
		for _, n := range neighbours {
			inRange := 0
			for _, b := range bots {
				if isInRange(n, b) {
					inRange++
				}
			}

			if inRange > currentInRange {
				//fmt.Printf("In range found in search %v at %v,%v,%v\n", inRange, n.x, n.y, n.z)
				best = n
				found = true
				currentInRange = inRange
			} else if inRange == currentInRange {
				start := nanobot{0, 0, 0, 0}
				if manhattanDistance(n, start) < manhattanDistance(pos, start) {
					//fmt.Printf("Found closer in search %v at %v,%v,%v\n", inRange, n.x, n.y, n.z)
					best = n
					found = true
					currentInRange = inRange
				}
			}
		}

		if !found {
			fmt.Printf("Found closer in search %v at %v,%v,%v\n", currentInRange, best.x, best.y, best.z)
			return best
		} else {
			pos = best
		}
	}
}

func getNeighbours(n nanobot) []nanobot {
	neighbours := make([]nanobot, 0)

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				neighbours = append(neighbours, nanobot{n.x + dx, n.y + dy, n.z + dz, 0})
			}
		}
	}

	return neighbours
}

func middle(i int, i2 int) int {
	middle := Abs(i-i2) / 2
	if i < i2 {
		return i + middle
	} else {
		return i2 + middle
	}
}

func evaluateArea(id int, minX int, maxX int, minY int, maxY int, minZ int, maxZ int, stepSize int, bots []nanobot) nanobot {
	maxInRange := math.MinInt64
	currentPos := nanobot{0, 0, 0, 0}
	start := nanobot{0, 0, 0, 0}
	loops := 0
	startTime := time.Now()
	total := (maxX - minX) * (maxY - minY) * (maxZ - minZ)
	for x := minX; x <= maxX; x += stepSize {
		for y := minY; y <= maxY; y += stepSize {
			for z := minZ; z <= maxZ; z += stepSize {
				pos := nanobot{x, y, z, 0}
				inRange := 0
				for _, b := range bots {
					if isInRange(pos, b) {
						inRange++
					}
				}

				if inRange == maxInRange {
					if manhattanDistance(start, pos) < manhattanDistance(start, currentPos) {
						currentPos = pos
						fmt.Printf("Pos closer, in range %v at %v,%v,%v\n", inRange, pos.x, pos.y, pos.z)
					}
				} else if inRange > maxInRange {
					currentPos = pos
					maxInRange = inRange
					fmt.Printf("In range updated %v at %v,%v,%v\n", inRange, pos.x, pos.y, pos.z)
				}
				loops++

				if loops%100000 == 0 {
					t := time.Now()
					elapsed := t.Sub(startTime)
					fmt.Printf("%v Progress %v in %v\n", id, float64(loops)/float64(total), elapsed)
				}
			}
		}
	}

	return currentPos
}

func botsInRangeOfStrongest(bots []nanobot) int {
	bot := bots[0]
	for _, b := range bots {
		if b.r > bot.r {
			bot = b
		}
	}
	inRange := 0
	for _, b := range bots {
		if isInRange(b, bot) {
			inRange++
		}
	}
	return inRange
}

func isInRange(target nanobot, strongest nanobot) bool {
	dist := manhattanDistance(target, strongest)
	return dist <= strongest.r
}

func manhattanDistance(target nanobot, strongest nanobot) int {
	return Abs(target.x-strongest.x) + Abs(target.y-strongest.y) + Abs(target.z-strongest.z)
}
