package day23

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"regexp"
)

type nanobot struct {
	x int
	y int
	z int
	r int
}

var ORIGO = nanobot{0, 0, 0, 0}

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

	maxX := 0
	maxY := 0
	maxZ := 0
	minX := 9999
	minY := 9999
	minZ := 9999

	for _, b := range bots {
		maxX = Max(b.x, maxX)
		maxY = Max(b.y, maxY)
		maxZ = Max(b.z, maxZ)
		minX = Min(b.x, minX)
		minY = Min(b.y, minY)
		minZ = Min(b.z, minZ)
	}

	size := 1
	curr := ORIGO
	max := Max(Max(maxX, maxY), maxZ)
	min := Min(Min(minX, minY), minZ)
	for {
		if size < (max - min) {
			size *= 2
		} else {
			break
		}
	}

	for {
		cubes := splitIntoCubes(curr, size)
		bestCube := cubes[0]
		bestCubeCount := -1
		for _, c := range cubes {
			count := 0
			for _, b := range bots {
				if isInRangeWithBoost(c, b, c.r) {
					count++
				}
			}
			if count > bestCubeCount {
				bestCube = c
				bestCubeCount = count
			} else if count == bestCubeCount {
				if manhattanDistance(c, ORIGO) < manhattanDistance(bestCube, ORIGO) {
					bestCube = c
					bestCubeCount = count
				}
			}
		}

		// fmt.Printf("Best cube %v %v\n", bestCube, bestCubeCount)
		if size == 1 {
			bestCube = checkNeighbours(bestCube, bestCubeCount, bots)
			return fmt.Sprint(manhattanDistance(bestCube, ORIGO))
		}

		size = size / 2
		curr = bestCube
	}
}

func checkNeighbours(curr nanobot, bestCount int, bots []nanobot) nanobot {
	c := 0
	for {
		neighbours := getNeighbours(curr)
		found := false
		for _, n := range neighbours {
			inRange := 0
			for _, b := range bots {
				if isInRange(n, b) {
					inRange++
				}
			}

			if inRange > c || inRange == c && manhattanDistance(n, ORIGO) < manhattanDistance(curr, ORIGO) {
				curr = n
				c = inRange
				found = true
			}
		}

		if !found {
			return curr
		}
	}
}

func splitIntoCubes(n nanobot, size int) []nanobot {
	splits := make([]nanobot, 0)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				splits = append(splits, nanobot{n.x + x*size, n.y + y*size, n.z + z*size, size})
			}
		}
	}
	return splits
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

func isInRangeWithBoost(target nanobot, strongest nanobot, boost int) bool {
	dist := manhattanDistance(target, strongest)
	return dist <= (strongest.r + boost)
}

func manhattanDistance(target nanobot, strongest nanobot) int {
	return Abs(target.x-strongest.x) + Abs(target.y-strongest.y) + Abs(target.z-strongest.z)
}
