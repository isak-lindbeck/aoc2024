package day04

import (
	. "github.com/isak-lindbeck/aoc2024/src/utils"
	"slices"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	directions := []Vector{
		North, South, // ↑ ↓
		West, East, // ← →
		SouthWest, NorthEast, // ↙ ↗
		NorthWest, SouthEast, // ↖ ↘
	}

	matrix := RuneMatrix(input)

	for x, y := range matrix.Keys() {
		value := matrix.Get(x, y)
		if value == 'X' {
			for direction := range slices.Values(directions) {
				word := []rune{'X', '.', '.', '.'}
				for index := 1; index < 4; index++ {
					d := direction.Mul(index)
					word[index] = matrix.GetSafe(x+d.X, y+d.Y, '.')
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
