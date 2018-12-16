package day03

import (
	"fmt"
	"github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"strconv"
	"strings"
)

func Solve(lines []string, partOne bool) string {

	claimedBy := make(map[pkg.Coordinate][]string, 0)
	claimsFree := make(map[string]bool, 0)
	for _, l := range lines {
		parts := strings.Split(l, " ")
		id := parts[0][1:]
		size := parts[2]
		pos := pkg.ParseCoordinate(size[:len(size)-1])
		dims := strings.Split(parts[3], "x")
		width, _ := strconv.Atoi(dims[0])
		height, _ := strconv.Atoi(dims[1])
		claimsFree[id] = true

		for x := pos.X; x < pos.X+width; x++ {
			for y := pos.Y; y < pos.Y+height; y++ {
				c := pkg.Coordinate{X: x, Y: y}
				claimed, f := claimedBy[c]
				if !f {
					claimed = make([]string, 0)
				}
				claimed = append(claimed, id)
				claimedBy[c] = claimed
			}
		}
	}

	moreThanTwo := 0

	for _, v := range claimedBy {
		if len(v) > 1 {
			moreThanTwo++

			for _, claimId := range v {
				delete(claimsFree, claimId)
			}
		}
	}

	if partOne {
		return fmt.Sprint(moreThanTwo)
	} else {
		for k := range claimsFree {
			return k
		}
	}

	return "Error"
}
