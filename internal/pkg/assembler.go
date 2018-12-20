package pkg

func Opr(registers []int, a int, b int, c int, op func(i1 int, i2 int) int) {
	registers[c] = op(registers[a], registers[b])
}

func Opi(registers []int, a int, b int, c int, op func(i1 int, i2 int) int) {
	registers[c] = op(registers[a], b)
}

func Addr(registers []int, a int, b int, c int) {
	Opr(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 + i2
	})
}

func Addi(registers []int, a int, b int, c int) {
	Opi(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 + i2
	})
}

func Mulr(registers []int, a int, b int, c int) {
	Opr(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 * i2
	})
}

func Muli(registers []int, a int, b int, c int) {
	Opi(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 * i2
	})
}

func Banr(registers []int, a int, b int, c int) {
	Opr(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 & i2
	})
}

func Bani(registers []int, a int, b int, c int) {
	Opi(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 & i2
	})
}

func Borr(registers []int, a int, b int, c int) {
	Opr(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 | i2
	})
}

func Bori(registers []int, a int, b int, c int) {
	Opi(registers, a, b, c, func(i1 int, i2 int) int {
		return i1 | i2
	})
}

func Setr(registers []int, a int, b int, c int) {
	registers[c] = registers[a]
}

func Seti(registers []int, a int, b int, c int) {
	registers[c] = a
}

func Gtir(registers []int, a int, b int, c int) {
	test := 0
	if a > registers[b] {
		test = 1
	}
	registers[c] = test
}

func Gtri(registers []int, a int, b int, c int) {
	test := 0
	if registers[a] > b {
		test = 1
	}
	registers[c] = test
}

func Gtrr(registers []int, a int, b int, c int) {
	test := 0
	if registers[a] > registers[b] {
		test = 1
	}
	registers[c] = test
}

func Eqir(registers []int, a int, b int, c int) {
	test := 0
	if a == registers[b] {
		test = 1
	}
	registers[c] = test
}

func Eqri(registers []int, a int, b int, c int) {
	test := 0
	if registers[a] == b {
		test = 1
	}
	registers[c] = test
}

func Eqrr(registers []int, a int, b int, c int) {
	test := 0
	if registers[a] == registers[b] {
		test = 1
	}
	registers[c] = test
}
