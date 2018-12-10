package day09

import (
	"fmt"
	"github.com/dr-horv/advent-of-code-2018/internal/pkg"
)

type marble struct {
	Value int
	Next  *marble
	Prev  *marble
}

func Solve(lines []string, partOne bool) string {
	curr := &marble{0, nil, nil}
	curr.Next = curr
	curr.Prev = curr
	player := 1
	numberOfPlayers := 418
	lastMarble := 71339
	if !partOne {
		lastMarble *= 100
	}
	scores := make(map[int]int, numberOfPlayers)
	//first := curr
	for i := 1; i <= lastMarble; i++ {

		if i%23 == 0 {
			r, n := remove(-7, curr)
			scores[player] = scores[player] + r.Value + i
			curr = n
		} else {
			newMarble := &marble{i, nil, nil}
			add(1, curr, newMarble)
			curr = newMarble
		}

		//printGame(first, curr, player)

		player = (player + 1) % numberOfPlayers
	}

	maxScore := 0
	for _, v := range scores {
		if v > maxScore {
			maxScore = v
		}
	}

	return fmt.Sprint(maxScore)
}

func printGame(first *marble, current *marble, player int) {
	info := fmt.Sprintf("[%v]", player+1)
	firstIteration := true
	m := first
	for {
		info += " "
		if !firstIteration && m.Value == first.Value {
			break
		}
		if current.Value == m.Value {
			info += fmt.Sprintf("(%v)", m.Value)
		} else {

			info += fmt.Sprintf(" %v ", m.Value)
		}

		m = m.Next
		firstIteration = false
	}
	fmt.Println(info)
}

func add(steps int, start *marble, element *marble) {
	curr := move(steps, start)
	next := curr.Next

	curr.Next = element
	element.Prev = curr
	element.Next = next
	next.Prev = element
}

func remove(steps int, start *marble) (*marble, *marble) {
	curr := move(steps, start)

	prev := curr.Prev
	next := curr.Next

	prev.Next = next
	next.Prev = prev

	return curr, next
}

func move(steps int, start *marble) *marble {
	move := func(m *marble) *marble {
		return m.Next
	}
	if steps < 0 {
		move = func(m *marble) *marble {
			return m.Prev
		}
	}
	curr := start
	steps = pkg.Abs(steps)
	for i := 0; i < steps; i++ {
		curr = move(curr)
	}
	return curr
}
