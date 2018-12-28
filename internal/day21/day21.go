package day21

import (
	"fmt"
	"github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"strings"
)

type op func(registers []int, a int, b int, c int)
type inststruction func(registers []int)

type ins struct {
	instr inststruction
	index int
	name  string
}

func Solve(lines []string, partOne bool) string {

	parts := strings.Split(lines[0], " ")
	pointerRegistry := pkg.StrConv(parts[1])
	program := make([]ins, 0)

	index := 0
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
		case "eqri":
			f = pkg.Eqri
		case "eqrr":
			f = pkg.Eqrr
		default:
			panic("Failed to parse: " + instruction)
		}
		fmt.Println(instruction)
		inst := func(registers []int) { f(registers, a, b, c) }
		inst(make([]int, 6))
		program = append(program, ins{inst, index, instruction})
		index++
	}

	fmt.Println("Parsing done")

	instructions := 0
	registers := make([]int, 6)
	registers[0] = 0
	ds := make(map[int]bool)
	lastD := -1

	for {
		nextInstructionIndex := registers[pointerRegistry]
		if nextInstructionIndex < 0 || nextInstructionIndex >= len(program) {
			fmt.Printf("Value: %v\n", registers[0])
			fmt.Printf("After: %v\n", registers)
			fmt.Printf("Instructions: %v\n", instructions)
			break
		}

		ins := program[nextInstructionIndex]

		if ins.index == 17 {
			times := registers[1] / 256
			registers[5] = times
			registers[4] = (times + 1) * 256
			registers[pointerRegistry] = 19

		} else {
			inst := ins.instr
			inst(registers)
		}

		if ins.index == 28 {
			d := registers[3]

			if partOne {
				return fmt.Sprint(d)
			}

			_, f := ds[d]
			if f {
				return fmt.Sprint(lastD)
			} else {
				ds[d] = true
				lastD = d
			}

		}

		registers[pointerRegistry] = registers[pointerRegistry] + 1
		instructions++

	}

	return "error"
}
