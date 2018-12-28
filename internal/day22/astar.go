package day22

import (
	"container/heap"
	"errors"
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"math"
)

type mapWithDefault struct {
	Map     map[State]int
	Default int
}

func (mwd mapWithDefault) get(s State) int {
	v, f := mwd.Map[s]

	if !f {
		return mwd.Default
	}

	return v
}

func (mwd mapWithDefault) set(s State, value int) {
	mwd.Map[s] = value
}

type QueueState struct {
	s State
	Index int
}

type Item struct {
	value    State
	priority int
	index 	int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type Path []State

func heuristic(s1 State, s2 State) int {
	dist := 0
	if s1.gear != s2.gear {
		dist += 7
	}

	return ManhattanDistance(s1.pos, s2.pos) + dist
}

func coordEquipmentKey(s State) string {
	return fmt.Sprintf("%v:%v", s.pos,s.gear)
}

func AStar(start State, goal State, world map[Coordinate]CavernType) (Path, error) {
	cameFrom := make(map[State]State)
	closedSet := make(map[string]int)
	openSlim := make(map[string]int)
	gScore := mapWithDefault{make(map[State]int), math.MaxInt64}
	gScore.set(start, 0)

	fScore := mapWithDefault{make(map[State]int), math.MaxInt64}
	fScore.set(start, ManhattanDistance(start.pos, goal.pos))

	priorityQueue := make(PriorityQueue, 1)
	priorityQueue[0] = &Item{
		start,
		fScore.get(start),
		0,
	}

	heap.Init(&priorityQueue)


	loop := 0
	for {
		if len(priorityQueue) == 0 {
			fmt.Println("No path")
			return nil, errors.New(fmt.Sprintf("no path found between %v amd %v", start, goal))
		}

		current:= heap.Pop(&priorityQueue).(*Item).value

		if loop % 10000 == 0 {
			fmt.Printf("Open set %v\n", len(priorityQueue))
			fmt.Printf("Closed set %v\n", len(closedSet))
			fmt.Printf("Currently exploring %v\n", current)
		}

		if current.pos == goal.pos && current.gear == goal.gear {
			fmt.Println("Found path")
			fmt.Printf("Open set %v\n", len(priorityQueue))
			fmt.Printf("Closed set %v\n", len(closedSet))
			path := ReconstructPath(cameFrom, current)
			//fmt.Printf("Start %v Goal %v, Path %v\n", start, goal, path)
			return path, nil
		}

		closedSet[coordEquipmentKey(current)] = current.timeSpent

		for _, neighbour := range GetNeighboursOfCaveState(current, world) {
			csf, found := closedSet[coordEquipmentKey(neighbour)]
			if (found && csf < neighbour.timeSpent) || neighbour.pos.X > 810 || neighbour.pos.Y > 810 {
				continue
			}

			tentativeGScore := gScore.get(current) + (neighbour.timeSpent - current.timeSpent)
			_, isOpen := openSlim[coordEquipmentKey(neighbour)]

			if !isOpen {
				openSlim[coordEquipmentKey(neighbour)] = neighbour.timeSpent
				priorityQueue.Push(&Item{
					value:neighbour,
					priority:tentativeGScore,
				})

			} else if tentativeGScore >= gScore.get(neighbour) {
				continue
			}

			cameFrom[neighbour] = current
			gScore.set(neighbour, tentativeGScore)
			fScore.set(neighbour, gScore.get(neighbour)+heuristic(neighbour, goal))
		}
		loop++
	}

}

func GetNeighboursOfCaveState(state State, world map[Coordinate]CavernType) []State {
	neighbours := make([]State, 0)
	ct := world[state.pos]
	possible := []Coordinate{state.pos.Up(), state.pos.Down(), state.pos.Left(), state.pos.Right()}

	for _, p := range possible {
		if p.X < 0 || p.Y < 0 {
			continue
		}

		nt, f := world[p]

		if !f {
			panic("Out of bounds: " + fmt.Sprint(p))
		}

		switch nt {
		case rocky:
			if state.gear == climbingGear || state.gear == torch {
				neighbours = append(neighbours, State{p, state.gear, state.timeSpent+1})
			}
		case wet:
			if state.gear == climbingGear || state.gear == neither {
				neighbours = append(neighbours, State{p, state.gear, state.timeSpent+1})
			}
		case narrow:
			if state.gear == torch || state.gear == neither {
				neighbours = append(neighbours, State{p, state.gear, state.timeSpent+1})
			}
		}
	}


	switch ct {
	case rocky:
		if state.gear == climbingGear {
			neighbours = append(neighbours, State{state.pos, torch, state.timeSpent+7})
		} else {
			neighbours = append(neighbours, State{state.pos, climbingGear, state.timeSpent+7})
		}
	case wet:
		if state.gear == climbingGear {
			neighbours = append(neighbours, State{state.pos, neither, state.timeSpent+7})
		} else {
			neighbours = append(neighbours, State{state.pos, climbingGear, state.timeSpent+7})
		}
	case narrow:
		if state.gear == torch {
			neighbours = append(neighbours, State{state.pos, neither, state.timeSpent+7})
		} else {
			neighbours = append(neighbours, State{state.pos, torch, state.timeSpent+7})
		}
	}

	return neighbours
}

func ReconstructPath(states map[State]State, curr State) []State {
	var path = []State{curr}
	for {
		cameFrom, found := states[curr]
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