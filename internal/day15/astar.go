package day15

import (
	"errors"
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"math"
)

type mapWithDefault struct {
	Map     map[Coordinate]int
	Default int
}

func (mwd mapWithDefault) get(coordinate Coordinate) int {
	v, f := mwd.Map[coordinate]

	if !f {
		return mwd.Default
	}

	return v
}

func (mwd mapWithDefault) set(coordinate Coordinate, value int) {
	mwd.Map[coordinate] = value
}

type Path []Coordinate

func AStar(start Coordinate, goal Coordinate, world map[Coordinate]Entity) (Path, error) {
	cameFrom := make(map[Coordinate]Coordinate)
	closedSet := make(map[Coordinate]bool)
	openSet := make(map[Coordinate]bool)
	openSet[start] = true

	gScore := mapWithDefault{make(map[Coordinate]int), math.MaxInt64}
	gScore.set(start, 0)

	fScore := mapWithDefault{make(map[Coordinate]int), math.MaxInt64}
	fScore.set(start, ManhattanDistance(start, goal))

	for {
		if len(openSet) == 0 {
			return nil, errors.New(fmt.Sprintf("no path found between %v amd %v", start, goal))
		}

		current := findMin(openSet, fScore)

		if current == goal {
			//fmt.Println("Found path")
			path := reconstructPath(cameFrom, current)
			//fmt.Printf("Start %v Goal %v, Path %v\n", start, goal, path)
			return path, nil
		}

		delete(openSet, current)
		closedSet[current] = true

		for _, neighbour := range getNeighbours(current, world) {
			_, found := closedSet[neighbour]
			if found {
				continue
			}

			tentativeGScore := gScore.get(current) + ManhattanDistance(current, neighbour)
			_, isOpen := openSet[neighbour]
			if !isOpen {
				openSet[neighbour] = true
			} else if tentativeGScore >= gScore.get(neighbour) {
				continue
			}

			cameFrom[neighbour] = current
			gScore.set(neighbour, tentativeGScore)
			fScore.set(neighbour, gScore.get(neighbour)+ManhattanDistance(neighbour, goal))
		}
	}

}

func getNeighbours(coordinate Coordinate, entities map[Coordinate]Entity) []Coordinate {
	neighbours := make([]Coordinate, 0)
	for _, c := range GetNeighboursSlice(coordinate) {
		entity := entities[c]
		if !entity.isObstruction() {
			neighbours = append(neighbours, c)
		}
	}

	return neighbours
}

func reconstructPath(coordinates map[Coordinate]Coordinate, curr Coordinate) []Coordinate {
	var path = []Coordinate{curr}
	for {
		cameFrom, found := coordinates[curr]
		if !found {
			break
		}

		path = append(path, cameFrom)
		curr = cameFrom
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func findMin(openSet map[Coordinate]bool, fScore mapWithDefault) Coordinate {
	minValue := math.MaxInt64
	var minC Coordinate
	for k := range openSet {
		value := fScore.get(k)

		if value < minValue {
			minValue = value
			minC = k
		} else if value == minValue && LessThanReadingOrder(k, minC) {
			minValue = value
			minC = k
		}

	}

	return minC
}
