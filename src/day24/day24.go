package day24

import (
	"fmt"
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"strings"
)

func Run(input string) (int, string) {
	ans1 := 0
	ans2 := ""

	split := strings.Split(input, "\n\n")

	state := make(map[string]int)
	for _, line := range strings.Split(split[0], "\n") {
		x := strings.Split(line, ": ")
		state[x[0]] = ints.Parse(x[1])
	}

	s := strings.Split(split[1], "\n")
	operations := make([]Operation, len(s))
	for i, line := range s {
		fields := strings.Fields(line)
		operations[i] = Operation{
			a:      fields[0],
			gateId: fields[1],
			gate:   op(fields[1]),
			b:      fields[2],
			target: fields[4],
		}

		//if operations[i].target == "z11" {
		//	operations[i].target = "vkq"
		//	continue
		//}
		//if operations[i].target == "vkq" {
		//	operations[i].target = "z11"
		//	continue
		//}

		if operations[i].target == "z24" {
			operations[i].target = "mmk"
			continue
		}
		if operations[i].target == "mmk" {
			operations[i].target = "z24"
			continue
		}

		if operations[i].target == "pvb" {
			operations[i].target = "qdq"
			continue
		}
		if operations[i].target == "qdq" {
			operations[i].target = "pvb"
			continue
		}

		if operations[i].target == "z38" {
			operations[i].target = "hqh"
			continue
		}
		if operations[i].target == "hqh" {
			operations[i].target = "z38"
			continue
		}
	}

	calculate(operations, state)

	for id, value := range state {
		if trimmed, hasPrefix := strings.CutPrefix(id, "z"); hasPrefix {
			ans1 += value << ints.Parse(trimmed)
		}
	}

	earliestWrongBit := getEarliestWrongBit(state, operations)
	fmt.Println(earliestWrongBit)
	fixAdderAtBit(operations, earliestWrongBit)
	return ans1, ans2
}

func fixAdderAtBit(operations []Operation, bit int) {
	z := fmt.Sprintf("z%02d", bit)
	x := fmt.Sprintf("x%02d", bit)
	//y := fmt.Sprintf("y%02d", bit)
	var xXorY Operation
	var xAndY Operation
	var xorToZ Operation
	for _, operation := range operations {
		if operation.target == z {
			xorToZ = operation
		}
		if operation.a == x || operation.b == x {
			if operation.gateId == "XOR" {
				xXorY = operation
			}
			if operation.gateId == "AND" {
				xAndY = operation
			}
		}
	}
	fmt.Println(xXorY)
	fmt.Println(xAndY)
	fmt.Println(xorToZ)

	fmt.Println(z)
}

func getOutput(state map[string]int) int {
	out := 0
	for id, value := range state {
		if trimmed, hasPrefix := strings.CutPrefix(id, "z"); hasPrefix {
			out += value << ints.Parse(trimmed)
		}
	}
	return out
}

func getEarliestWrongBit(state map[string]int, operations []Operation) int {
	earliest := 0
	for i := 0; earliest == 0 && i < 45; i++ {
		for j := 0; earliest == 0 && j < 45; j++ {
			x := 1<<i - 1
			y := 1<<j - 1
			setStartState(&state, x, y)
			calculate(operations, state)
			out := getOutput(state)
			expected := x + y
			if out != expected {
				foundFirstWrongBit := 0
				for foundFirstWrongBit != 1 {
					out = out >> 1
					expected = expected >> 1
					foundFirstWrongBit = (out & 1) ^ (expected & 1)
					earliest++
				}
			}
		}
	}
	return earliest
}

func calculate(operations []Operation, state map[string]int) {
	idx := 0
	size := len(operations)
	for size > 0 {
		i := idx % size
		o := operations[i]

		_, aExists := state[o.a]
		_, bExists := state[o.b]
		if aExists && bExists {
			res := o.gate(state[o.a], state[o.b])
			state[o.target] = res
			operations[i], operations[size-1] = operations[size-1], operations[i]
			size--
		}

		idx++
	}
}

func setStartState(state *map[string]int, x, y int) {
	for id, _ := range *state {
		if trimmed, hasPrefix := strings.CutPrefix(id, "x"); hasPrefix {
			(*state)[id] = (x >> ints.Parse(trimmed)) & 1
			continue
		}
		if trimmed, hasPrefix := strings.CutPrefix(id, "y"); hasPrefix {
			(*state)[id] = (y >> ints.Parse(trimmed)) & 1
			continue
		}
		delete(*state, id)

	}
}

type Operation struct {
	a         string
	gateId    string
	gate      func(a int, b int) int
	b, target string
}

func op(s string) func(a int, b int) int {
	switch s {
	case "AND":
		return and
	case "OR":
		return or
	case "XOR":
		return xor
	}
	panic("invalid operation: " + s)
}

func and(a, b int) int {
	return a & b
}

func or(a, b int) int {
	return a | b
}

func xor(a, b int) int {
	return a ^ b
}
