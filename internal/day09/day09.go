package day09

import (
	"fmt"
	"time"
)

func Solve(lines []string, partOne bool) string {

	start := time.Now()
	curr := 1
	circle := make([]int, 2)
	circle[0] = 0
	circle[1] = 1
	player := 1
	numberOfPlayers := 418
	lastMarble := 71339 * 100
	scores := make(map[int]int, numberOfPlayers)
	for i := 2; i <= lastMarble; i++ {

		if i%23 == 0 {
			indexToRemove := getIndex(curr-7, circle)
			e := circle[indexToRemove]
			scores[player] = scores[player] + e + i
			circle = removeElement(indexToRemove, circle)
			curr = indexToRemove
		} else {
			placeIndex := getIndex(curr+1, circle) + 1
			circle = addElement(placeIndex, i, circle)
			if placeIndex <= curr {
				curr++
			}

			if getIndex(placeIndex+2, circle) == curr || getIndex(placeIndex-2, circle) == curr {
				curr = placeIndex
			}
		}

		//printGame(circle, curr, player)
		if i%100000 == 0 {
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Printf("Progress %v in %v\n", float64(i)/float64(lastMarble), elapsed)
		}
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

func printGame(circle []int, current int, player int) {
	info := fmt.Sprintf("[%v]", player+1)
	for i, n := range circle {
		info += " "
		if current == i {
			info += fmt.Sprintf("(%v)", n)
		} else {

			info += fmt.Sprintf(" %v ", n)
		}
	}
	fmt.Println(info)
}

func getIndex(index int, slice []int) int {
	return (index + len(slice)) % len(slice)
}

func removeElement(index int, slice []int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func addElement(index int, element int, slice []int) []int {
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = element

	return slice
}
