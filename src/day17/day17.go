package day17

import (
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"regexp"
	"strconv"
	"strings"
)

func Run(input string) (string, int) {
	ans1 := ""
	ans2 := 0

	re := regexp.MustCompile("[-0-9]+")

	allNumbers := re.FindAllString(input, -1)

	operands := make([]int, 0)
	for i := 3; i < len(allNumbers); i++ {
		operands = append(operands, ints.Parse(allNumbers[i]))
	}

	program := Program{
		regA: ints.Parse(allNumbers[0]),
		regB: ints.Parse(allNumbers[1]),
		regC: ints.Parse(allNumbers[2]),
		code: operands}
	out := runProgram(program)
	var stringSlice []string
	for _, i := range out {
		stringSlice = append(stringSlice, strconv.Itoa(i))
	}
	ans1 = strings.Join(stringSlice, ",")

	foundBase := make([]int, len(operands))

	for exp := len(operands) - 1; exp >= 0; exp-- {
		cycle := ints.Pow(8, exp)
		base := 0
		for i := len(operands) - 1; i > exp; i-- {
			base += foundBase[i]
		}
		found := false
		for i := foundBase[exp]; i < (cycle * 8); i += (cycle) {
			program = Program{
				regA: i + base,
				regB: ints.Parse(allNumbers[1]),
				regC: ints.Parse(allNumbers[2]),
				code: operands}
			out = runProgram(program)

			sameLen := len(out) == len(operands)
			if sameLen && out[exp] == operands[exp] {
				foundBase[exp] = i
				found = true
				break
			}
		}
		if !found {
			foundBase[exp] = 0
			foundBase[exp+1] += ints.Pow(8, exp+1)
			exp += 2
		}
	}
	for _, v := range foundBase {
		ans2 += v
	}

	return ans1, ans2
}

func runProgram(program Program) []int {
	out := make([]int, 0)
	idx := 0
	for idx < len(program.code) {
		op := program.code[idx]
		idx++
		switch op {
		case 0:
			program.regA = program.regA / ints.Pow(2, program.combo(idx))
			idx++
		case 1:
			program.regB = program.regB ^ program.literal(idx)
			idx++
		case 2:
			program.regB = program.combo(idx) % 8
			idx++
		case 3:
			if program.regA != 0 {
				idx = program.literal(idx)
			} else {
				idx++
			}
		case 4:
			program.regB = program.regB ^ program.regC
			idx++
		case 5:
			out = append(out, program.combo(idx)%8)
			idx++
		case 6:
			program.regB = program.regA / ints.Pow(2, program.combo(idx))
			idx++
		case 7:
			program.regC = program.regA / ints.Pow(2, program.combo(idx))
			idx++
		}
	}
	return out
}

type Program struct {
	regA, regB, regC int
	code             []int
}

func (p *Program) literal(idx int) int {
	return p.code[idx]
}

func (p *Program) combo(idx int) int {
	i := p.code[idx]
	if i <= 3 {
		return i
	}

	switch i {
	case 4:
		return p.regA
	case 5:
		return p.regB
	case 6:
		return p.regC
	}
	panic("Unknown opcode" + strconv.Itoa(i))
}
