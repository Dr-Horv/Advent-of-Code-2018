package pkg

import (
	"math"
	"strconv"
	"strings"
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

func Sum(slice []int) int {
	s := 0
	for _, v := range slice {
		s += v
	}

	return s
}

func StrConv(s string) int {
	v, e := strconv.Atoi(strings.TrimSpace(s))
	if e != nil {
		panic("Can't parse " + s)
	}

	return v
}
