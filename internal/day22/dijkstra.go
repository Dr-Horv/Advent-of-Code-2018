package day22

import (
	"errors"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"math"
)



func Dijkstra(start State, goal State, world map[Coordinate]CavernType) (Path, error) {

	unvisited := make(map[State]bool)
	tentativeDistances := mapWithDefault{make(map[State]int), math.MaxInt64}
	tentativeDistances.set(start, 0)
	cameFrom := make(map[State]State)

	current := start
	for {

		if current.pos == goal.pos && current.gear == goal.gear {
			return ReconstructPath(cameFrom, current), nil
		}

		for _, n := range GetNeighboursOfCaveState(current, world) {
			d := tentativeDistances.get(n)
			if d > n.timeSpent {
				tentativeDistances.set(n, n.timeSpent)
			}

			unvisited[n] = true
			cameFrom[n] = current
		}
		delete(unvisited, current)

		for k, _ := range unvisited {
			current = k
			break
		}

		if len(unvisited) == 0 {
			return nil, errors.New("no path")
		}




	}





}