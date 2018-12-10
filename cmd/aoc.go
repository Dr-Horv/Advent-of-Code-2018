package main

import (
	"fmt"
	"github.com/dr-horv/advent-of-code-2018/internal/day10"
	"github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"time"
)

func main() {
	var day = "10"
	lines := pkg.ReadFile("./internal/day" + day + "/input")
	start := time.Now()
	answer := day10.Solve(lines, false)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(answer)
	fmt.Println(elapsed)
}
