package day10

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type signal struct {
	Position Coordinate
	DX       int
	DY       int
}

func Solve(lines []string, partOne bool) string {
	var re = regexp.MustCompile(`(?m)position=<(.*)> velocity=<(.*)>`)
	signals := make([]*signal, 0)

	for _, l := range lines {
		groups := re.FindStringSubmatch(l)

		x, _ := strconv.Atoi(strings.TrimSpace(strings.Split(groups[1], ",")[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(strings.Split(groups[1], ",")[1]))
		dx, _ := strconv.Atoi(strings.TrimSpace(strings.Split(groups[2], ",")[0]))
		dy, _ := strconv.Atoi(strings.TrimSpace(strings.Split(groups[2], ",")[1]))

		signals = append(signals, &signal{Coordinate{X: x, Y: y}, dx, dy})
	}

	distance := math.MaxFloat64
	loops := 0
	prints := 0
	lastSky := ""
	for {
		newDistance := measureDistance(signals)
		if newDistance > distance {
			break
		}
		distance = newDistance
		if distance < float64(30) {
			lastSky = getSky(signals)
			prints++
		}

		for _, s := range signals {
			s.Position = Coordinate{X: s.Position.X + s.DX, Y: s.Position.Y + s.DY}
		}
		loops++
	}

	return fmt.Sprintf("%vSpotted after %v seconds", lastSky, loops-1)
}

func measureDistance(signals []*signal) float64 {

	sum := 0
	count := 0
	for _, s1 := range signals {
		for _, s2 := range signals {
			sum += ManhattanDistance(s1.Position, s2.Position)
			count++
		}
	}

	return float64(sum) / float64(count)

}

func getSky(signals []*signal) string {
	x0 := math.MaxInt64
	y0 := math.MaxInt64
	x1 := 0
	y1 := 0
	signalMap := make(map[Coordinate]bool, len(signals))

	for _, s := range signals {
		c := s.Position
		x0 = Min(c.X, x0)
		x1 = Max(c.X, x1)

		y0 = Min(c.X, y0)
		y1 = Max(c.Y, y1)
		signalMap[Coordinate{X: c.X, Y: c.Y}] = true
	}

	var skyBuilder strings.Builder
	topLeft := Coordinate{X: x0, Y: y0}
	bottomRight := Coordinate{X: x1, Y: y1}

	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		for x := topLeft.X; x <= bottomRight.X; x++ {
			_, found := signalMap[Coordinate{x, y}]
			if found {
				skyBuilder.WriteString("#")
			} else {
				skyBuilder.WriteString(" ")
			}
		}
		skyBuilder.WriteString("\n")
	}

	return skyBuilder.String()
}
