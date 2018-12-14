package pkg

import (
	"math"
)

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func Compare(i1 int, i2 int, operator func(int, int) bool) int {
	if operator(i1, i2) {
		return i1
	} else {
		return i2
	}
}

func Min(i1 int, i2 int) int {
	return Compare(i1, i2, func(i1 int, i int) bool {
		return i1 < i2
	})
}

func Max(i1 int, i2 int) int {
	return Compare(i1, i2, func(i1 int, i int) bool {
		return i1 > i2
	})
}

func FindMax(slice []int) (int, int) {
	max := math.MinInt64
	maxI := -1
	for i, v := range slice {
		if v > max {
			max = v
			maxI = i
		}
	}

	return max, maxI
}
