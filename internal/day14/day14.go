package day14

import (
	"fmt"
	"math"
	"strconv"
)

func Solve(lines []string, partOne bool) string {
	// input := 360781

	recipes := make([]int, 2)
	recipes[0] = 3
	recipes[1] = 7
	elfOneIndex := 0
	elfTwoIndex := 1

	target, _ := strconv.Atoi(lines[0])
	numbers := 6
	targetSlice := make([]int, 0)
	for i := 0; i < numbers; i++ {
		v := target / int(math.Pow10(6-i-1))
		targetSlice = append(targetSlice, v%10)
	}

	fmt.Println(targetSlice)
	scanIndex := 0
	for {
		sum := recipes[elfOneIndex] + recipes[elfTwoIndex]
		loops := 1
		if sum >= 10 && sum < 20 {
			loops = 2
		}

		for j := 0; j < loops; j++ {
			curr := sum / int(math.Pow10(loops-j-1))
			recipes = append(recipes, curr%10)
		}

		elfOneIndex = (elfOneIndex + 1 + recipes[elfOneIndex]) % len(recipes)
		elfTwoIndex = (elfTwoIndex + 1 + recipes[elfTwoIndex]) % len(recipes)
		if partOne {
			if len(recipes) >= target+10 {

				scoreNextTen := ""
				for i := 0; i < 10; i++ {
					scoreNextTen += fmt.Sprint(recipes[target+i])
				}

				return scoreNextTen
			}
		} else {
			if len(recipes) < scanIndex+len(targetSlice) {
				continue
			}
			test := true
			for i := 0; i < len(targetSlice); i++ {
				if recipes[scanIndex+i] != targetSlice[i] {
					test = false
					break
				}
			}
			if test {
				return fmt.Sprint(scanIndex)
			} else {
				scanIndex++
			}

		}
	}
}

func prettyPrint(recipes []int, elf1 int, elf2 int) {
	str := ""
	for i := 0; i < len(recipes); i++ {
		if i == elf1 {
			str += "(" + fmt.Sprint(recipes[i]) + ")"
		} else if i == elf2 {
			str += "[" + fmt.Sprint(recipes[i]) + "]"
		} else {
			str += " " + fmt.Sprint(recipes[i]) + " "
		}
	}

	fmt.Println(str)

}
