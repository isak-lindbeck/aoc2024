package day25

import (
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"strings"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	locks := make([][]int, 0)
	keys := make([][]int, 0)

	for _, text := range strings.Split(input, "\n\n") {
		isLock := strings.HasPrefix(text, "#####")
		matrix := utils.RuneMatrix(text)
		code := make([]int, 5)
		for x := 0; x < matrix.Width; x++ {
			for y := 1; y < matrix.Height-1; y++ {
				if matrix.Get(x, y) == '#' {
					code[x]++
				}
			}
		}
		if isLock {
			locks = append(locks, code)
		} else {
			keys = append(keys, code)
		}

	}

	for _, lock := range locks {
		for _, key := range keys {
			match := 1
			for i := 0; i < 5; i++ {
				if lock[i]+key[i] > 5 {
					match = 0
				}
			}
			ans1 += match
		}
	}

	return ans1, ans2
}
