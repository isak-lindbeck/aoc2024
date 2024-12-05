package day05

import (
	"fmt"
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"slices"
	"strings"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	input = strings.TrimSuffix(input, "\n")

	split := strings.Split(input, "\n\n")
	pairs := strings.Split(split[0], "\n")

	var greaterThanMap = make(map[GreaterThan]bool)
	for p := range slices.Values(pairs) {
		pInts := strings.Split(p, "|")
		gt := GreaterThan{high: pInts[0], low: pInts[1]}
		greaterThanMap[gt] = true
	}
	fmt.Println(split[0])
	fmt.Println(split[1])

	lists := strings.Split(split[1], "\n")
	for list := range slices.Values(lists) {
		l := strings.Split(list, ",")
		isSorted := slices.IsSortedFunc(l, sortByMap(greaterThanMap))
		if isSorted {
			ans1 += ints.Parse(l[(len(l) / 2)])
		} else {
			slices.SortFunc(l, sortByMap(greaterThanMap))
			ans2 += ints.Parse(l[(len(l) / 2)])
		}
	}

	return ans1, ans2
}

type GreaterThan struct {
	high string
	low  string
}

func sortByMap(gtMap map[GreaterThan]bool) func(a string, b string) int {
	return func(a, b string) int {
		if gtMap[GreaterThan{high: a, low: b}] {
			return -1
		}
		return 1
	}
}
