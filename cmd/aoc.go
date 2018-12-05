package main

import (
	"fmt"
	"github.com/Dr-Horv/Advent-of-Code-2018/internal/day05"
	"github.com/Dr-Horv/Advent-of-Code-2018/internal/pkg"
	"time"
)

func main() {
	var day = "05"
	lines := pkg.ReadFile("./internal/day" + day + "/input")
	start := time.Now()
	answer := day05.Solve(lines, false)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(answer)
	fmt.Println(elapsed)
}
