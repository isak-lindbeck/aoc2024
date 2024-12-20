package day19

import (
	"strings"
	"sync"
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

	var wg sync.WaitGroup
	ch := make(chan (int), len(patterns))

	for _, pattern := range patterns {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- calcPossibleCombinations(pattern, towels, make(map[string]int))
		}()
	}
	wg.Wait()
	close(ch)
	for i := range ch {
		ans2 += i
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
