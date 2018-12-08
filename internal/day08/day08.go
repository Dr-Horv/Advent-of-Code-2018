package day08

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	ID               int
	NumberOfChildren int
	NumberOfMetadata int
	MetadataEntries  []int
	Children         []node
}

func Solve(lines []string, partOne bool) string {
	input := lines[0]
	numbers := parseInput(input)
	i := 0
	root, _, _ := parseNode(0, i, numbers)

	sum := 0
	if partOne {
		sum = summarizeMetadataIncludingChildren(root)
	} else {
		sum = summarizeValueOf(root)
	}

	return fmt.Sprint(sum)

}

func summarizeValueOf(n node) int {
	if n.NumberOfChildren == 0 {
		return summarizeMetadata(n)
	}

	sum := 0
	for i := 0; i < n.NumberOfMetadata; i++ {
		ci := n.MetadataEntries[i] - 1
		if ci >= 0 && ci < n.NumberOfChildren {
			sum += summarizeValueOf(n.Children[ci])
		}
	}

	return sum
}

func summarizeMetadataIncludingChildren(n node) int {
	sum := summarizeMetadata(n)
	for i := 0; i < n.NumberOfChildren; i++ {
		sum += summarizeMetadataIncludingChildren(n.Children[i])
	}

	return sum
}

func summarizeMetadata(n node) int {
	sum := 0
	for i := 0; i < n.NumberOfMetadata; i++ {
		sum += n.MetadataEntries[i]
	}
	return sum
}

func parseNode(id int, index int, data []int) (node, int, int) {
	n := node{id, data[index], data[index+1], make([]int, 0), make([]node, 0)}
	id++
	newIndex := index + 2
	var cn node
	for ci := 0; ci < n.NumberOfChildren; ci++ {
		cn, newIndex, id = parseNode(id, newIndex, data)
		n.Children = append(n.Children, cn)
	}

	for di := 0; di < n.NumberOfMetadata; di++ {
		n.MetadataEntries = append(n.MetadataEntries, data[newIndex])
		newIndex++
	}

	return n, newIndex, id
}

func parseInput(s string) []int {
	numbers := make([]int, 0)

	for _, p := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(strings.TrimSpace(p))
		numbers = append(numbers, i)
	}

	return numbers
}
