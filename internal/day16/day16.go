package day16

import (
	"fmt"
	"strconv"
	"strings"
)

func opr(registers []int, a int, b int, c int, op func(i1 int, i2 int) int) {
	registers[c] = op(registers[a], registers[b])
}

func opi(registers []int, a int, b int, c int, op func(i1 int, i2 int) int) {
	registers[c] = op(registers[a], b)
}

func addr(registers []int, a int, b int, c int) {
	opr(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 + i2
	})
}

func addi(registers []int, a int, b int, c int) {
	opi(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 + i2
	})
}

func mulr(registers []int, a int, b int, c int) {
	opr(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 * i2
	})
}

func muli(registers []int, a int, b int, c int) {
	opi(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 * i2
	})
}

func banr(registers []int, a int, b int, c int) {
	opr(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 & i2
	})
}

func bani(registers []int, a int, b int, c int) {
	opi(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 & i2
	})
}

func borr(registers []int, a int, b int, c int) {
	opr(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 | i2
	})
}

func bori(registers []int, a int, b int, c int) {
	opi(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 | i2
	})
}

func setr(registers []int, a int, b int, c int) {
	registers[c] = registers[a]
}

func seti(registers []int, a int, b int, c int) {
	registers[c] = a
}

func gtir(registers []int, a int, b int, c int) {
	test := 0
	if a > registers[b] {
		test = 1
	}
	registers[c] = test
}

func gtri(registers []int, a int, b int, c int) {
	test := 0
	if registers[a] > b {
		test = 1
	}
	registers[c] = test
}

func gtrr(registers []int, a int, b int, c int) {
	test := 0
	if registers[a] > registers[b] {
		test = 1
	}
	registers[c] = test
}

func eqir(registers []int, a int, b int, c int) {
	test := 0
	if a == registers[b] {
		test = 1
	}
	registers[c] = test
}

func eqri(registers []int, a int, b int, c int) {
	test := 0
	if registers[a] == b {
		test = 1
	}
	registers[c] = test
}

func eqrr(registers []int, a int, b int, c int) {
	test := 0
	if registers[a] == registers[b] {
		test = 1
	}
	registers[c] = test
}

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

func strConv(s string) int {
	v, e := strconv.Atoi(strings.TrimSpace(s))
	if e != nil {
		panic("Can't parse " + s)
	}

	return v
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
				origRegisters[0] = strConv(values[0])
				origRegisters[1] = strConv(values[1])
				origRegisters[2] = strConv(values[2])
				origRegisters[3] = strConv(values[3])
			} else if strings.Contains(l, "After") {
				data := l[len("After:  [") : len(l)-1]
				values := strings.Split(data, ",")
				expectedRegisters = make([]int, 4)
				expectedRegisters[0] = strConv(values[0])
				expectedRegisters[1] = strConv(values[1])
				expectedRegisters[2] = strConv(values[2])
				expectedRegisters[3] = strConv(values[3])

				samples = append(samples, sample{origRegisters, expectedRegisters, id, a, b, c})
			} else {
				values := strings.Split(strings.TrimSpace(l), " ")
				id = strConv(values[0])
				a = strConv(values[1])
				b = strConv(values[2])
				c = strConv(values[3])
			}
		} else {
			values := strings.Split(strings.TrimSpace(l), " ")
			program = append(program, programInstruction{
				strConv(values[0]),
				strConv(values[1]),
				strConv(values[2]),
				strConv(values[3]),
			})

		}
	}

	operations := []*operation{
		{addi},
		{addr},
		{mulr},
		{muli},
		{banr},
		{bani},
		{borr},
		{bori},
		{setr},
		{seti},
		{gtir},
		{gtri},
		{gtrr},
		{eqir},
		{eqri},
		{eqrr}}

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
