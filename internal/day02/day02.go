package day02

import "fmt"

func Solve(lines []string, partOne bool) string {

	if partOne {
		return doPartOne(lines)
	} else {
		return doPartTwo(lines)
	}
}

func doPartTwo(lines []string) string {

	for li, l1 := range lines {
		for lj, l2 := range lines {
			if li == lj {
				continue
			}

			test := true
			errorFound := false
			for i := 0; i < len(l2); i++ {
				if l1[i] != l2[i] {
					if errorFound {
						test = false
						break
					}
					errorFound = true
				}
			}

			if test {
				commonLetters := ""
				for i := 0; i < len(l1); i++ {
					if l1[i] == l2[i] {
						commonLetters += string(l1[i])
					}
				}
				return commonLetters
			}

		}
	}

	return "Error"
}

func doPartOne(lines []string) string {
	threes := 0
	twos := 0
	for _, l := range lines {

		runeCount := make(map[rune]int, 0)

		for _, r := range l {
			c, found := runeCount[r]
			if !found {
				c = 0
			}
			c++
			runeCount[r] = c
		}

		hasTwoLetter := false
		hasThreeLetter := false

		for _, rc := range runeCount {
			if rc == 3 {
				hasThreeLetter = true
			} else if rc == 2 {
				hasTwoLetter = true
			}
		}

		if hasTwoLetter {
			twos++
		}

		if hasThreeLetter {
			threes++
		}

	}
	return fmt.Sprint(twos * threes)
}
