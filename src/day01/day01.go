package day01

import (
	"sort"
	"strconv"
	"strings"
)

func Run(input string) (int, int) {
	input = strings.TrimSuffix(input, "\n")

	var sliceA = make([]int, 1)
	var sliceB = make([]int, 1)
	var mapB = make(map[int]int)
	for _, line := range strings.Split(input, "\n") {
		res := strings.Fields(line)
		aValue, _ := strconv.Atoi(res[0])
		bValue, _ := strconv.Atoi(res[1])

		sliceA = append(sliceA, aValue)
		sliceB = append(sliceB, bValue)

		mapB[bValue] += 1
	}

	sort.Ints(sliceA)
	sort.Ints(sliceB)

	var result1 = 0
	var result2 = 0
	for i, aVal := range sliceA {
		bVal := sliceB[i]
		diff := absDiffInt(aVal, bVal)
		result1 += diff
		result2 += aVal * mapB[aVal]
	}

	return result1, result2
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
