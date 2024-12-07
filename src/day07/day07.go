package day07

import (
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"strconv"
	"strings"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	input = strings.TrimSuffix(input, "\n")
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, ":")
		goalValue := ints.Parse(split[0])
		fields := strings.Fields(split[1])
		parts := make([]int, len(fields))
		for i, s := range fields {
			parts[i] = ints.Parse(s)
		}
		if CalculatePart1(goalValue, len(parts), parts) {
			ans1 += goalValue
		}
		if CalculatePart2(goalValue, len(parts), parts) {
			ans2 += goalValue
		}
	}

	return ans1, ans2
}

func CalculatePart1(goalValue, partIndex int, parts []int) bool {
	partIndex--
	if partIndex < 0 {
		return goalValue == 0
	}

	part := parts[partIndex]

	res := CalculatePart1(goalValue-part, partIndex, parts)
	if goalValue%part == 0 {
		res = res || CalculatePart1(goalValue/part, partIndex, parts)
	}

	return res
}

func CalculatePart2(goalValue, partIndex int, parts []int) bool {
	partIndex--
	if partIndex < 0 {
		return goalValue == 0
	}

	part := parts[partIndex]

	res := CalculatePart2(goalValue-part, partIndex, parts)
	if goalValue%part == 0 {
		res = res || CalculatePart2(goalValue/part, partIndex, parts)
	}

	pow := ints.Pow(10, len(strconv.Itoa(part)))
	if goalValue%pow == part {
		res = res || CalculatePart2(goalValue/pow, partIndex, parts)
	}

	return res
}
