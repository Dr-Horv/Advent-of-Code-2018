package main

import (
	"fmt"
	"github.com/dr-horv/advent-of-code-2018/internal/day09"
	"github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"time"
)

func main() {
	var day = "09"
	lines := pkg.ReadFile("./internal/day" + day + "/input")
	start := time.Now()
	answer := day09.Solve(lines, false)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(answer)
	fmt.Println(elapsed)
}
