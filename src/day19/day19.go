package day19

import (
	"strings"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	split := strings.Split(input, "\n\n")

	towels := strings.Split(split[0], ", ")
	patterns := strings.Split(split[1], "\n")

	for _, pattern := range patterns {
		isPossible := calcIsPossible(pattern, towels)
		if isPossible {
			ans1++
		}
	}

	for _, pattern := range patterns {
		possibleCombinations := calcPossibleCombinations(pattern, towels, make(map[string]int))
		ans2 += possibleCombinations
	}

	return ans1, ans2
}

func calcIsPossible(pattern string, towels []string) bool {
	for _, towel := range towels {
		if after, found := strings.CutSuffix(pattern, towel); found {
			if after == "" || calcIsPossible(after, towels) {
				return true
			}
		}
	}
	return false
}

func calcPossibleCombinations(pattern string, towels []string, cache map[string]int) int {
	if value, exists := cache[pattern]; exists {
		return value
	}

	sum := 0
	for _, towel := range towels {
		if after, found := strings.CutSuffix(pattern, towel); found {
			if after == "" {
				sum++
			} else {
				sum += calcPossibleCombinations(after, towels, cache)
			}
		}
	}
	cache[pattern] = sum
	return sum
}
