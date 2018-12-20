package day16

import (
	"fmt"
	"github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"strings"
)

type sample struct {
	Registers []int
	Expected  []int
	ID        int
	A         int
	B         int
	C         int
}

func (sample sample) copyRegisters() []int {
	r := make([]int, 4)
	r[0] = sample.Registers[0]
	r[1] = sample.Registers[1]
	r[2] = sample.Registers[2]
	r[3] = sample.Registers[3]

	return r
}

func (sample sample) copyExpectedRegisters() []int {
	r := make([]int, 4)
	r[0] = sample.Expected[0]
	r[1] = sample.Expected[1]
	r[2] = sample.Expected[2]
	r[3] = sample.Expected[3]

	return r
}

type instruction func(registers []int, a int, b int, c int)
type operation struct {
	instruction instruction
}

type programInstruction struct {
	ID int
	A  int
	B  int
	C  int
}

func Solve(lines []string, partOne bool) string {
	origRegisters := make([]int, 4)
	expectedRegisters := make([]int, 4)
	id := -1
	samples := make([]sample, 0)
	program := make([]programInstruction, 0)
	var a int
	var b int
	var c int
	for i, l := range lines {
		if len(strings.TrimSpace(l)) == 0 {
			continue
		}

		if i <= 3093 {
			if strings.Contains(l, "Before") {
				data := l[len("Before: [") : len(l)-1]
				values := strings.Split(data, ",")
				origRegisters = make([]int, 4)
				origRegisters[0] = pkg.StrConv(values[0])
				origRegisters[1] = pkg.StrConv(values[1])
				origRegisters[2] = pkg.StrConv(values[2])
				origRegisters[3] = pkg.StrConv(values[3])
			} else if strings.Contains(l, "After") {
				data := l[len("After:  [") : len(l)-1]
				values := strings.Split(data, ",")
				expectedRegisters = make([]int, 4)
				expectedRegisters[0] = pkg.StrConv(values[0])
				expectedRegisters[1] = pkg.StrConv(values[1])
				expectedRegisters[2] = pkg.StrConv(values[2])
				expectedRegisters[3] = pkg.StrConv(values[3])

				samples = append(samples, sample{origRegisters, expectedRegisters, id, a, b, c})
			} else {
				values := strings.Split(strings.TrimSpace(l), " ")
				id = pkg.StrConv(values[0])
				a = pkg.StrConv(values[1])
				b = pkg.StrConv(values[2])
				c = pkg.StrConv(values[3])
			}
		} else {
			values := strings.Split(strings.TrimSpace(l), " ")
			program = append(program, programInstruction{
				pkg.StrConv(values[0]),
				pkg.StrConv(values[1]),
				pkg.StrConv(values[2]),
				pkg.StrConv(values[3]),
			})

		}
	}

	operations := []*operation{
		{pkg.Addi},
		{pkg.Addr},
		{pkg.Mulr},
		{pkg.Muli},
		{pkg.Banr},
		{pkg.Bani},
		{pkg.Borr},
		{pkg.Bori},
		{pkg.Setr},
		{pkg.Seti},
		{pkg.Gtir},
		{pkg.Gtri},
		{pkg.Gtrr},
		{pkg.Eqir},
		{pkg.Eqri},
		{pkg.Eqrr}}

	behaveLikeThree := 0
	opMap := make(map[int]*operation)
	foundOps := make(map[*operation]bool)

	for {
		for _, sample := range samples {
			_, foundOpAlready := opMap[sample.ID]
			if !partOne && foundOpAlready {
				continue
			}
			tests := 0
			expectedRegisters := sample.copyExpectedRegisters()
			var potential *operation
			for _, op := range operations {
				_, found := foundOps[op]
				if !partOne && found {
					continue
				}
				registers := sample.copyRegisters()
				f := op.instruction
				f(registers, sample.A, sample.B, sample.C)

				if registers[0] == expectedRegisters[0] &&
					registers[1] == expectedRegisters[1] &&
					registers[2] == expectedRegisters[2] &&
					registers[3] == expectedRegisters[3] {
					tests++
					potential = op
				}
			}
			if tests >= 3 {
				behaveLikeThree++
			}

			if tests == 1 {
				opMap[sample.ID] = potential
				foundOps[potential] = true
			}
		}

		if partOne {
			break
		}
		if len(opMap) == 16 {
			break
		}
	}

	if partOne {
		return fmt.Sprint(behaveLikeThree)
	}

	registers := make([]int, 4)
	registers[0] = 0
	registers[1] = 0
	registers[2] = 0
	registers[3] = 0

	for _, pi := range program {
		op := opMap[pi.ID]
		op.instruction(registers, pi.A, pi.B, pi.C)
	}

	return fmt.Sprint(registers[0])
}
