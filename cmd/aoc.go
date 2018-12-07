package main

import (
	"fmt"
	"github.com/dr-horv/advent-of-code-2018/internal/day07"
	"github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"time"
)

func main() {
	var day = "07"
	lines := pkg.ReadFile("./internal/day" + day + "/input")
	start := time.Now()
	answer := day07.Solve(lines, false)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(answer)
	fmt.Println(elapsed)
}
