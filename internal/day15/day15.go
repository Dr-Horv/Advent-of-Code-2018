package day15

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
)

type elf struct {
	ID          string
	Pos         Coordinate
	AttackPower int
	HP          int
}

func (e elf) isObstruction() bool {
	return true
}
func (e elf) String() string {
	return "E"
}

func (e elf) isElf() bool {
	return true
}

func (e elf) isGoblin() bool {
	return false
}

type goblin struct {
	ID          string
	Pos         Coordinate
	AttackPower int
	HP          int
}

func (g goblin) isObstruction() bool {
	return true
}

func (g goblin) String() string {
	return "G"
}

func (e goblin) isElf() bool {
	return false
}

func (e goblin) isGoblin() bool {
	return true
}

type empty rune

func (e empty) isObstruction() bool {
	return false
}

func (e empty) String() string {
	return "."
}

func (e empty) isElf() bool {
	return false
}

func (e empty) isGoblin() bool {
	return false
}

type wall rune

func (w wall) isObstruction() bool {
	return true
}

func (w wall) String() string {
	return "#"
}

func (w wall) isElf() bool {
	return false
}

func (w wall) isGoblin() bool {
	return false
}

type Entity interface {
	isObstruction() bool
	String() string
	isElf() bool
	isGoblin() bool
}

func LessThanReadingOrder(c1 Coordinate, c2 Coordinate) bool {
	if c1.Y < c2.Y {
		return true
	}

	if c1.Y == c2.Y && c1.X < c2.X {
		return true
	}

	return false
}

func Solve(lines []string, partOne bool) string {
	origCaveMap := make(map[Coordinate]Entity)
	origGoblinMap := make(map[Coordinate]goblin)
	origElfMap := make(map[Coordinate]elf)
	height := len(lines)
	width := len(lines[0])
	id := 'a'
	elves := 0
	for y, l := range lines {
		for x, r := range l {
			c := Coordinate{X: x, Y: y}
			var e Entity
			if r == 'G' {
				goblin := goblin{string(id), c, 3, 200}
				id++
				origGoblinMap[c] = goblin
				e = goblin
			} else if r == 'E' {
				elf := elf{string(id), c, 3, 200}
				id++
				origElfMap[c] = elf
				elves++
				e = elf
			} else if r == '#' {
				e = wall(r)
			} else if r == '.' {
				e = empty(r)
			}
			origCaveMap[c] = e
		}
	}

	attackPower := 3
	var hasDoneTurn map[string]bool
	var round int
	for {
		round = 0
		caveMap := make(map[Coordinate]Entity, len(origCaveMap))
		for k, v := range origCaveMap {
			caveMap[k] = v
		}
		goblinMap := make(map[Coordinate]goblin, len(origGoblinMap))
		for k, v := range origGoblinMap {
			goblinMap[k] = v
		}
		elfMap := make(map[Coordinate]elf, len(origElfMap))
		for k, v := range origElfMap {
			v.AttackPower = attackPower
			elfMap[k] = v
		}

		fmt.Printf("Starting simulation with %v attack power\n", attackPower)
		//prettyPrintMap(caveMap, elfMap, goblinMap, width, height)

		for {
			hasDoneTurn = make(map[string]bool, 0)
			if !partOne {
				if len(elfMap) < elves {
					break
				}
			}
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					c := Coordinate{X: x, Y: y}
					e, _ := caveMap[c]

					if e.isElf() {
						elf := elfMap[c]
						_, hasTurned := hasDoneTurn[elf.ID]
						if hasTurned {
							continue
						}
						doElfTurn(elfMap, c, goblinMap, caveMap)
						hasDoneTurn[elf.ID] = true
					} else if e.isGoblin() {
						goblin := goblinMap[c]
						_, hasTurned := hasDoneTurn[goblin.ID]
						if hasTurned {
							continue
						}
						doGoblinTurn(elfMap, c, goblinMap, caveMap)
						hasDoneTurn[goblin.ID] = true
					}
				}
			}
			//fmt.Printf("Round %v\n", round)
			//prettyPrintMap(caveMap, elfMap, goblinMap, width, height)

			if len(elfMap) == 0 || len(goblinMap) == 0 {
				fmt.Printf("End at round: %v\n", round)
				prettyPrintMap(caveMap, elfMap, goblinMap, width, height)
				return calculateAnswer(round, elfMap, goblinMap)
			}

			round++

		}

		attackPower++

	}

}

func calculateAnswer(round int, elves map[Coordinate]elf, goblins map[Coordinate]goblin) string {
	sum := 0
	if len(elves) == 0 {
		for _, g := range goblins {
			sum += g.HP
		}
	} else {
		for _, e := range elves {
			sum += e.HP
		}
	}

	return fmt.Sprint(round * sum)
}

func doGoblinTurn(elfMap map[Coordinate]elf, c Coordinate, goblinMap map[Coordinate]goblin, caveMap map[Coordinate]Entity) {
	goblin := goblinMap[c]
	//fmt.Printf("Turn for %v\n", goblin.ID)
	targetPositions := make([]Coordinate, 0)
	targetElves := make([]elf, 0)
	for ep, e := range elfMap {
		dist := ManhattanDistance(ep, goblin.Pos)
		if dist == 1 {
			targetElves = append(targetElves, e)
		}

		for _, ptg := range GetNeighboursSlice(ep) {
			if !caveMap[ptg].isObstruction() {
				targetPositions = append(targetPositions, ptg)
			}
		}
	}

	if len(targetElves) == 0 {
		var closest Path
		for _, tp := range targetPositions {
			for _, n := range GetNeighboursSlice(goblin.Pos) {
				if caveMap[n].isObstruction() {
					continue
				}

				path, err := AStar(n, tp, caveMap)

				if err != nil {
					continue
				}
				if closest == nil || len(path) < len(closest) {
					closest = path
				} else if len(path) == len(closest) {
					if LessThanReadingOrder(path[0], closest[0]) {
						closest = path
					}

				}
			}
		}

		if closest != nil {
			// fmt.Printf("Found path for %v %v moving to %v\n", goblin.ID, goblin.Pos, closest[0])
			caveMap[goblin.Pos] = empty('.')
			delete(goblinMap, goblin.Pos)
			goblin.Pos = closest[0]
			goblinMap[goblin.Pos] = goblin
			caveMap[goblin.Pos] = goblin

			for _, n := range GetNeighboursSlice(goblin.Pos) {
				e, found := elfMap[n]

				if found {
					targetElves = append(targetElves, e)
				}
			}
		}

	}

	if len(targetElves) > 0 {
		minHp := targetElves[0].HP
		minHpG := targetElves[0]

		for _, g := range targetElves {
			if g.HP < minHp {
				minHp = g.HP
				minHpG = g
			} else if g.HP == minHp && LessThanReadingOrder(g.Pos, minHpG.Pos) {
				minHp = g.HP
				minHpG = g
			}
		}

		minHpG.HP -= goblin.AttackPower

		if minHpG.HP <= 0 {
			delete(elfMap, minHpG.Pos)
			caveMap[minHpG.Pos] = empty('.')
		} else {
			caveMap[minHpG.Pos] = minHpG
			elfMap[minHpG.Pos] = minHpG
		}

	}

}

func doElfTurn(elfMap map[Coordinate]elf, c Coordinate, goblinMap map[Coordinate]goblin, caveMap map[Coordinate]Entity) {
	elf := elfMap[c]
	targetPositions := make([]Coordinate, 0)
	targetGoblins := make([]goblin, 0)
	for gp, g := range goblinMap {
		if ManhattanDistance(gp, elf.Pos) == 1 {
			targetGoblins = append(targetGoblins, g)
		}

		for _, ptg := range GetNeighboursSlice(gp) {
			if !caveMap[ptg].isObstruction() {
				targetPositions = append(targetPositions, ptg)
			}
		}
	}

	if len(targetGoblins) == 0 {
		var closest Path
		for _, tp := range targetPositions {
			for _, n := range GetNeighboursSlice(elf.Pos) {
				if caveMap[n].isObstruction() {
					continue
				}

				path, err := AStar(n, tp, caveMap)

				if err != nil {
					continue
				}
				if closest == nil || len(path) < len(closest) {
					closest = path
				} else if len(path) == len(closest) {
					if LessThanReadingOrder(path[0], closest[0]) {
						closest = path
					}
				}
			}
		}

		if closest != nil {
			caveMap[elf.Pos] = empty('.')
			delete(elfMap, elf.Pos)
			elf.Pos = closest[0]
			elfMap[elf.Pos] = elf
			caveMap[elf.Pos] = elf

			for _, n := range GetNeighboursSlice(elf.Pos) {
				g, found := goblinMap[n]

				if found {
					targetGoblins = append(targetGoblins, g)
				}
			}
		}

	}

	if len(targetGoblins) > 0 {
		minHp := targetGoblins[0].HP
		minHpG := targetGoblins[0]

		for _, g := range targetGoblins {
			if g.HP < minHp {
				minHp = g.HP
				minHpG = g
			} else if g.HP == minHp && LessThanReadingOrder(g.Pos, minHpG.Pos) {
				minHp = g.HP
				minHpG = g
			}
		}

		minHpG.HP -= elf.AttackPower

		if minHpG.HP <= 0 {
			delete(goblinMap, minHpG.Pos)
			caveMap[minHpG.Pos] = empty('.')
		} else {
			caveMap[minHpG.Pos] = minHpG
			goblinMap[minHpG.Pos] = minHpG
		}

	}

}

func prettyPrintMap(entities map[Coordinate]Entity, elfMap map[Coordinate]elf, goblinMap map[Coordinate]goblin, width int, height int) {
	str := ""
	nbrOfUnits := 0

	for _, e := range entities {
		if e.isGoblin() || e.isElf() {
			nbrOfUnits++
		}
	}

	for y := 0; y < height; y++ {
		units := make([]string, 0)
		for x := 0; x < width; x++ {
			c := Coordinate{X: x, Y: y}
			e, f := entities[c]

			if !f {
				str += " "
				continue
			}

			if e.isElf() {
				elf := elfMap[c]
				units = append(units, fmt.Sprintf("E(%v)", elf.HP))
			}

			if e.isGoblin() {
				g := goblinMap[c]
				units = append(units, fmt.Sprintf("G(%v)", g.HP))
			}

			str += e.String()
		}

		str += "  "
		for _, u := range units {
			str += u + ", "
		}

		str += "\n"
	}

	fmt.Println(str)
}
