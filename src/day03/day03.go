package day03

import (
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"regexp"
)

func Run(input string) (int, int) {
	compile := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	allMatches := compile.FindAllString(input, -1)
	compile = regexp.MustCompile("\\d+")

	result1 := 0
	result2 := 0
	enabled := true
	for i := 0; i < len(allMatches); i++ {
		match := allMatches[i]
		if match == "do()" {
			enabled = true
			continue
		}
		if match == "don't()" {
			enabled = false
			continue
		}

		numStrings := compile.FindAllString(match, 2)
		a := ints.Parse(numStrings[0])
		b := ints.Parse(numStrings[1])
		result1 += a * b

		if enabled {
			result2 += a * b
		}
	}

	return result1, result2
}
