package day19

import (
	"fmt"
	"github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"strings"
)

type op func(registers []int, a int, b int, c int)
type inststruction func(registers []int)

func Solve(lines []string, partOne bool) string {

	parts := strings.Split(lines[0], " ")
	pointerRegistry := pkg.StrConv(parts[1])
	program := make([]inststruction, 0)

	for _, l := range lines[1:] {
		parts := strings.Split(l, " ")
		instruction := parts[0]
		a := pkg.StrConv(parts[1])
		b := pkg.StrConv(parts[2])
		c := pkg.StrConv(parts[3])

		var f op
		switch instruction {
		case "addr":
			f = pkg.Addr
		case "addi":
			f = pkg.Addi
		case "mulr":
			f = pkg.Mulr
		case "muli":
			f = pkg.Muli
		case "banr":
			f = pkg.Banr
		case "bani":
			f = pkg.Bani
		case "borr":
			f = pkg.Borr
		case "bori":
			f = pkg.Bori
		case "setr":
			f = pkg.Setr
		case "seti":
			f = pkg.Seti
		case "gtir":
			f = pkg.Gtir
		case "gtri":
			f = pkg.Gtri
		case "gtrr":
			f = pkg.Gtrr
		case "eqir":
			f = pkg.Eqir
		case "eqrr":
			f = pkg.Eqrr
		}
		fmt.Println(instruction)
		inst := func(registers []int) { f(registers, a, b, c) }
		inst(make([]int, 6))
		program = append(program, inst)
	}

	fmt.Println("Parsing done")
	registers := make([]int, 6)

	registers[0] = 1

	for {
		nextInstructionIndex := registers[pointerRegistry]
		if nextInstructionIndex < 0 || nextInstructionIndex >= len(program) {
			break
		}

		inst := program[nextInstructionIndex]
		//fmt.Printf("Before: %v\n", registers)
		inst(registers)
		//fmt.Printf("After: %v\n", registers)
		registers[pointerRegistry] = registers[pointerRegistry] + 1
	}

	fmt.Printf("After: %v\n", registers)

	return fmt.Sprint(registers[0])
}
