package day01

import (
	"fmt"
	"strconv"
)

func Solve(lines []string, partOne bool) string {

	numbers := make([]int, 0)
	sum := 0
	for _, l := range lines {
		i, _ := strconv.Atoi(l)
		sum += i
		numbers = append(numbers, i)
	}

	if partOne {
		return fmt.Sprint(sum)
	} else {
		frequency := 0
		frequencies := make(map[int]bool)
		frequencies[frequency] = true

		for {
			for _, n := range numbers {
				frequency += n

				_, found := frequencies[frequency]
				if found {
					return fmt.Sprint(frequency)
				}

				frequencies[frequency] = true
			}

		}

	}
}
