package day05

import (
	"fmt"
	"math"
	"unicode"
)

type job func(rune) string

func worker(id int, jobs <-chan rune, results chan<- string, job job) {
	for j := range jobs {
		result := job(j)
		results <- result
	}
}

func Solve(lines []string, partOne bool) string {

	input := lines[0]
	if partOne {
		s := process(input)
		return fmt.Sprint(len(s))
	}

	unitTypes := make(map[rune]bool)
	messages := make(chan string, 100)
	jobs := make(chan rune, 100)
	totalCandidates := 0

	task := func(upper rune) string {
		lower := unicode.ToLower(upper)
		newCandidate := getNewCandidate(input, upper, lower)
		stable := process(newCandidate)
		return stable
	}

	for w := 1; w <= 4; w++ {
		go worker(w, jobs, messages, task)
	}

	for _, r := range []rune(input) {
		_, found := unitTypes[r]
		if !found {
			totalCandidates++
			upperR := unicode.ToUpper(r)
			lowerR := unicode.ToLower(r)
			unitTypes[upperR] = true
			unitTypes[lowerR] = true
			jobs <- upperR
		}
	}
	close(jobs)

	min := math.MaxInt64
	minS := ""

	finished := 0
	for totalCandidates > finished {
		// fmt.Printf("Finished %v of %v total %v \n", finished, totalCandidates, float64(finished)/float64(totalCandidates))
		stable := <-messages
		if len(stable) < min {
			min = len(stable)
			minS = stable
		}
		finished++
	}

	return fmt.Sprint(len(minS))
}

func getNewCandidate(input string, upper rune, lower rune) string {
	s := []rune(input)
	var next []rune
	for i := 0; i < len(s); i++ {
		if s[i] == upper || s[i] == lower {
			continue
		}
		next = append(next, s[i])
	}

	return string(next)
}

func process(input string) string {
	runes := []rune(input)
	for i := 0; i < (len(runes) - 2); {
		// compare if two runes have different case i.e. a and A or B and b.
		if runes[i]^runes[i+1] == 32 {
			runes = append(runes[:i], runes[i+2:]...)
			if i > 0 {
				i--
			}
			continue
		}
		i++
	}

	return string(runes)
}
