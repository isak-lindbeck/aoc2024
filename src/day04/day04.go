package day04

import (
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"slices"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	directions := []Vector{
		{0, 1}, {0, -1}, // ↑ ↓
		{-1, 0}, {1, 0}, // ← →
		{-1, -1}, {1, 1}, // ↙ ↗
		{-1, 1}, {1, -1}, // ↖ ↘
	}

	matrix := utils.RuneMatrix(input)

	for x, y := range matrix.Keys() {
		value := matrix.Get(x, y)
		if value == 'X' {
			for direction := range slices.Values(directions) {
				word := []rune{'X', '.', '.', '.'}
				for index := 1; index < 4; index++ {
					d := direction.mul(index)
					word[index] = matrix.GetSafe(x+d.x, y+d.y, '.')
				}
				wordString := string(word)
				if wordString == "XMAS" {
					ans1++
				}
			}
		}
		if value == 'A' {
			word1 := string([]rune{matrix.GetSafe(x+1, y+1, '.'), 'A', matrix.GetSafe(x-1, y-1, '.')})
			word2 := string([]rune{matrix.GetSafe(x-1, y+1, '.'), 'A', matrix.GetSafe(x+1, y-1, '.')})
			if (word1 == "MAS" || word1 == "SAM") && (word2 == "MAS" || word2 == "SAM") {
				ans2++
			}
		}

	}

	return ans1, ans2
}

type Vector struct{ x, y int }

func (d Vector) mul(a int) Vector {
	return Vector{d.x * a, d.y * a}
}
