package day24

import (
	"fmt"
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"sort"
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
			left:   fields[0],
			gateId: fields[1],
			right:  fields[2],
			target: fields[4],
		}
	}

	runOperations(operations, state)
	ans1 = getOutputZValue(state)

	swapped := make([]string, 0)
	earliestWrongBit := 0
	for {
		earliestWrongBit = getEarliestWrongBit(state, operations, earliestWrongBit)
		if earliestWrongBit > 0 {
			swap1, swap2 := findAndSwapTargets(operations, earliestWrongBit)
			swapped = append(swapped, swap1, swap2)
		} else {
			break
		}
	}
	sort.Strings(swapped)
	ans2 = strings.Join(swapped, ",")
	return ans1, ans2
}

func runOperations(operations []Operation, state map[string]int) {
	idx := 0
	size := len(operations)
	for size > 0 {
		i := idx % size
		o := operations[i]

		_, leftExists := state[o.left]
		_, rightExists := state[o.right]
		if leftExists && rightExists {
			res := toGate(o.gateId)(state[o.left], state[o.right])
			state[o.target] = res
			operations[i], operations[size-1] = operations[size-1], operations[i]
			size--
		}
		idx++
	}
}

func getEarliestWrongBit(state map[string]int, operations []Operation, fromBit int) int {
	earliestWrong := 0
	for i := fromBit; earliestWrong == 0 && i < 45; i++ {
		x := 1<<i - 1
		y := 0
		setStartState(&state, x, y)
		runOperations(operations, state)
		out := getOutputZValue(state)
		expected := x + y
		if out != expected {
			foundFirstWrongBit := 0
			for foundFirstWrongBit != 1 {
				out = out >> 1
				expected = expected >> 1
				foundFirstWrongBit = (out & 1) ^ (expected & 1)
				earliestWrong++
			}
		}
	}
	return earliestWrong
}

func findAndSwapTargets(operations []Operation, wrongBit int) (string, string) {
	// op1: x XOR y -> var1
	// op2: in XOR var1 -> z
	// op3: x AND y -> var2
	// op4: in AND var1 -> var3
	// op5: var3 OR var2 -> out

	z := fmt.Sprintf("z%02d", wrongBit)
	x := fmt.Sprintf("x%02d", wrongBit)
	prevX := fmt.Sprintf("x%02d", wrongBit-1)
	prevOp3 := findByOperandGate(operations, prevX, "AND")
	prevOp5 := findByOperandGate(operations, prevOp3.target, "OR")
	in := prevOp5.target

	expectedOp2 := findByOperandGate(operations, in, "XOR")
	actualOp2 := findByTarget(operations, z)
	if expectedOp2 != actualOp2 {
		expectedOp2.target, actualOp2.target = actualOp2.target, expectedOp2.target
		return expectedOp2.target, actualOp2.target
	}
	op2 := expectedOp2

	var var1 string
	if op2.left == in {
		var1 = op2.right
	} else {
		var1 = op2.left
	}

	expectedOp1 := findByTarget(operations, var1)
	actualOp1 := findByOperandGate(operations, x, "XOR")
	if expectedOp1 != actualOp1 {
		expectedOp1.target, actualOp1.target = actualOp1.target, expectedOp1.target
		return expectedOp1.target, actualOp1.target
	}

	panic("could not fix error in adder")
}

func findByOperandGate(operations []Operation, operand, gateId string) *Operation {
	var found *Operation
	for i, operation := range operations {
		if operation.gateId == gateId {
			if operation.left == operand || operation.right == operand {
				found = &operations[i]
				break
			}
		}
	}
	return found
}

func findByTarget(operations []Operation, target string) *Operation {
	var found *Operation
	for i, operation := range operations {
		if operation.target == target {
			found = &operations[i]
			break
		}
	}
	return found
}

func getOutputZValue(state map[string]int) int {
	out := 0
	for id, value := range state {
		if trimmed, hasPrefix := strings.CutPrefix(id, "z"); hasPrefix {
			out += value << ints.Parse(trimmed)
		}
	}
	return out
}

func setStartState(state *map[string]int, x, y int) {
	for id := range *state {
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
	left   string
	gateId string
	right  string
	target string
}

func toGate(s string) func(a int, b int) int {
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
