package day12

import (
	"fmt"
	"strings"
	"time"
)

type plant struct {
	Index int
	Value bool
	Left  *plant
	Right *plant
}

//type rule func(bool, bool, bool, bool, bool) (matches bool, value bool)
type rule func(v1 bool, v2 bool, v3 bool, v4 bool, v5 bool) (matches bool, value bool)

func Solve(lines []string, partOne bool) string {
	initialState := "##..##....#.#.####........##.#.#####.##..#.#..#.#...##.#####.###.##...#....##....#..###.#...#.#.#.#"
	first := &plant{0, parseRune(rune(initialState[0])), nil, nil}
	root := padLeft(first, 3)
	curr := first
	for i, c := range initialState[1:] {
		newPlant := &plant{1 + i, parseRune(rune(c)), curr, nil}
		curr.Right = newPlant
		curr = newPlant
	}
	padRight(curr, 3)

	fmt.Println("Padding done")

	rules := make([]rule, 0)
	for _, l := range lines {
		rules = append(rules, parseRule(l))
	}

	first = root
	generations := 50000000000
	start := time.Now()
	for g := 0; g < generations; g++ {
		//printPlants(first, g)

		if g%1000 == 0 {
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Printf("Progress: %v after %v \n", float64(g)/float64(generations), elapsed)
		}

		first, last := nextGeneration(first, rules)
		if first.Value {
			first = padLeft(first, 1)
		}

		if last.Value || last.Left.Value || last.Left.Left.Value {
			last = padRight(last, 1)
		}
	}

	//printPlants(first, 20)

	sum := calculateSum(first)

	return fmt.Sprint(sum)
}

func calculateSum(curr *plant) int {
	sum := 0
	for {
		if curr == nil {
			break
		}

		if curr.Value {
			// fmt.Printf("Index %v\n", curr.Index)
			sum += curr.Index
		}

		curr = curr.Right
	}

	return sum
}

func parseRule(s string) rule {
	// "...## => #"
	parts := strings.Split(s, " => ")
	lhs := parts[0]
	rhs := parts[1]
	b1 := parseRune(rune(lhs[0]))
	b2 := parseRune(rune(lhs[1]))
	b3 := parseRune(rune(lhs[2]))
	b4 := parseRune(rune(lhs[3]))
	b5 := parseRune(rune(lhs[4]))
	v := parseRune(rune(rhs[0]))
	// fmt.Printf("%v %v %v %v %v gives %v\n", b1,b2,b3,b4,b5, v)

	/**
	...## => #
	..#.. => #
	.#... => #
	.#.#. => #
	.#.## => #
	.##.. => #
	.#### => #
	#.#.# => #
	#.### => #
	##.#. => #
	##.## => #
	###.. => #
	###.# => #
	####. => #
	*/

	r := func(v1 bool, v2 bool, v3 bool, v4 bool, v5 bool) (bool, bool) {

		if (b1 == v1) && (b2 == v2) && (b3 == v3) && (b4 == v4) && (b5 == v5) {
			return v, true
		}
		return false, false
	}

	// v, m := r(false, false, false, true, true)
	// fmt.Printf("match %v\n", m)

	return r
}

func nextGeneration(root *plant, rules []rule) (*plant, *plant) {
	curr := root.Right.Right
	l2v := curr.Left.Left.Value
	l1v := curr.Left.Value
	cv := curr.Value
	r1v := curr.Right.Value
	r2v := curr.Right.Right.Value
	for {
		for _, r := range rules {

			v, match := r(l2v, l1v, cv, r1v, r2v)
			if match {
				curr.Value = v
				//printPlants(root)
				break
			}
			curr.Value = false
		}

		curr = curr.Right
		if curr.Right.Right == nil {
			break
		}

		l2v = l1v
		l1v = cv
		cv = r1v
		r1v = r2v
		r2v = curr.Right.Right.Value

	}

	return root, curr.Right
}

func padLeft(curr *plant, size int) *plant {
	for i := 0; i < size; i++ {
		newPlant := &plant{curr.Index - 1, false, nil, nil}
		newPlant.Right = curr
		curr.Left = newPlant
		curr = newPlant
	}
	return curr
}

func padRight(curr *plant, size int) *plant {
	for i := 0; i < size; i++ {
		newPlant := &plant{curr.Index + 1, false, nil, nil}
		newPlant.Left = curr
		curr.Right = newPlant
		curr = newPlant
	}

	return curr
}

func printPlants(curr *plant, prefix int) {
	str := ""
	if prefix > 9 {
		str += fmt.Sprint(prefix)
	} else {
		str += " " + fmt.Sprint(prefix)
	}

	str += ": ."

	for {

		if curr.Value {
			str += "#"
		} else {
			str += "."
		}

		if curr.Right == nil {
			break
		}

		curr = curr.Right
	}

	str += "."
	fmt.Println(str)
}

func parseRune(rune rune) bool {
	return rune == '#'
}
