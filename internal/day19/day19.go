package day19

import (
	"fmt"
	"github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"math"
	"strings"
)

type op func(registers []int, a int, b int, c int)
type inststruction func(registers []int)

type ins struct {
	instr inststruction
	index int
}

func Solve(lines []string, partOne bool) string {

	if partOne {
		return doPartOne(lines)
	}

	a := 1 // 0
	b := 0 // 1
	c := 0 // 2
	d := 0 // 3
	e := 0 // 4
	f := 0 // 5, program pointer

	f = f + 16
	f++
	a, b, c, d, e, f = doLine17(a, b, c, d, e, f)
	f++
	if a == 0 {
		a, b, c, d, e, f = doLine1(a, b, c, d, e, f)
		return fmt.Sprint(a)
	} else {
		d = 27
		f = 27
		f++
		d = d * f
		f++
		d = d + f
		f++
		d = d * f
		f++
		d = d * 14
		f++
		d = d * f
		f++
		e = e + d
		a = 0
		f = 0
		a, b, c, d, e, f = doLine1(a, b, c, d, e, f)
		return fmt.Sprint(a)

	}

}

func doPartOne(lines []string) string {
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
		case "eqrr":
			f = pkg.Eqrr
		}
		inst := func(registers []int) { f(registers, a, b, c) }
		inst(make([]int, 6))
		program = append(program, ins{inst, index})
		index++
	}

	fmt.Println("Parsing done")
	registers := make([]int, 6)
	instructionsDone := 0
	for {
		nextInstructionIndex := registers[pointerRegistry]
		if nextInstructionIndex < 0 || nextInstructionIndex >= len(program) {
			return fmt.Sprint(registers[0])
		}

		ins := program[nextInstructionIndex]
		inst := ins.instr
		inst(registers)
		registers[pointerRegistry] = registers[pointerRegistry] + 1
		instructionsDone++
	}
}

func doLine17(a int, b int, c int, d int, e int, f int) (int, int, int, int, int, int) {
	e = e + 2
	f++
	e = e * e
	f++
	e = f * e
	f++
	e = e * 11
	f++
	d = d + 2
	f++
	d = d * f
	f++
	d = d + 13
	f++
	e = e + d
	f++
	return a, b, c, d, e, f
}

func doLine1(a int, b int, c int, d int, e int, f int) (int, int, int, int, int, int) {
	c = 1
	a = a + 1 + e
	for i := 2; i <= int(math.Sqrt(float64(e))); i++ {
		if e%i == 0 {
			a += i
			a += e / i
		}
	}
	return a, b, c, d, e, f
}
