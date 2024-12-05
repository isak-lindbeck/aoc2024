package day02

import (
	"cmp"
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"slices"
	"strings"
)

func Run(input string) (int, int) {
	input = strings.TrimSuffix(input, "\n")
	var zeroBad = 0
	var oneBad = 0
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		values := make([]int, len(fields))
		for i, s := range fields {
			values[i] = ints.Parse(s)
		}

		if checkIsGood(values) {
			zeroBad++
		} else {
			for i := 0; i < len(values); i++ {
				if checkIfGoodSkipIndex(values, i) {
					oneBad++
					break
				}
			}
		}
	}

	return zeroBad, zeroBad + oneBad
}

func checkIsGood(values []int) bool {
	return checkIfGoodSkipIndex(values, -1)
}

func checkIfGoodSkipIndex(values []int, skipIndex int) bool {
	if skipIndex >= 0 {
		values = removeIndex(values, skipIndex)
	}

	isWithinBounds := checkIncrementBounds(values)
	isSorted := slices.IsSorted(values) || slices.IsSortedFunc(values, reverse())
	return isWithinBounds && isSorted
}

func removeIndex(slice []int, index int) []int {
	newSlice := make([]int, len(slice)-1)
	offset := 0
	for i := range slice {
		if i == index {
			offset = 1
			continue
		}
		newSlice[i-offset] = slice[i]
	}
	return newSlice
}

func checkIncrementBounds(values []int) bool {
	for i := 0; i < len(values)-1; i++ {
		diff := values[i+1] - values[i]
		if ints.Abs(diff) < 1 || ints.Abs(diff) > 3 {
			return false
		}
	}
	return true
}

func reverse() func(a int, b int) int {
	return func(a, b int) int {
		return cmp.Compare(b, a)
	}
}
