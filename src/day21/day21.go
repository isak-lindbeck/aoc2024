package day21

import (
	"cmp"
	"github.com/isak-lindbeck/aoc2024/src/ints"
	. "github.com/isak-lindbeck/aoc2024/src/utils"
	"math"
	"slices"
	"strings"
)

var directions = []Vector{Up, Right, Down, Left}

var numpad = RuneMatrix("789\n456\n123\n#0A")
var keypad = RuneMatrix("#^A\n<v>")

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	split := strings.Split(input, "\n")

	for _, numCode := range split {
		runes := []rune("A" + numCode)
		codeNumValue := ints.Parse(strings.TrimSuffix(numCode, "A"))

		for i := 0; i < len(runes)-1; i++ {
			ans1 += SolveNumpad(runes[i], runes[i+1], 2) * codeNumValue
			ans2 += SolveNumpad(runes[i], runes[i+1], 25) * codeNumValue
		}
	}

	return ans1, ans2
}

func SolveNumpad(fromKey rune, toKey rune, keypadRobots int) int {
	from := NewVector(numpad.GetCoordinates(fromKey))
	to := NewVector(numpad.GetCoordinates(toKey))
	paths := SolvePaths(numpad, from, to)

	shortest := math.MaxInt
	for _, path := range paths {
		runes := []rune("A" + path)
		keyCode := 0
		for i := 0; i < len(runes)-1; i++ {
			keyCode += SolveKeypad(runes[i], runes[i+1], keypadRobots-1)
		}
		if keyCode < shortest {
			shortest = keyCode
		}
	}

	return shortest
}

type CacheKey struct {
	fromKey, toKey rune
	depth          int
}

var cache = make(map[CacheKey]int)

func SolveKeypad(fromKey rune, toKey rune, depth int) int {
	key := CacheKey{fromKey, toKey, depth}
	c := cache
	if value, exist := c[key]; exist {
		return value
	}
	from := NewVector(keypad.GetCoordinates(fromKey))
	to := NewVector(keypad.GetCoordinates(toKey))
	paths := SolvePaths(keypad, from, to)

	if depth == 0 {
		shortest := slices.MinFunc(paths, shortestLength())
		c[key] = len(shortest)
		return len(shortest)
	}

	shortest := math.MaxInt
	for _, path := range paths {
		runes := []rune("A" + path)
		keyCode := 0
		for i := 0; i < len(runes)-1; i++ {
			keyCode += SolveKeypad(runes[i], runes[i+1], depth-1)
		}
		if keyCode < shortest {
			shortest = keyCode
		}
	}
	c[key] = shortest
	return shortest
}

func SolvePaths(matrix Matrix[rune], from, target Vector) []string {
	out := make([]string, 0)

	if from == target {
		return []string{"A"}
	}

	for _, direction := range directions {
		next := from.Add(direction.XY())
		if next.Distance(target) >= from.Distance(target) {
			continue
		}

		if matrix.GetSafeAt(next.XY())('#') != '#' {
			children := SolvePaths(matrix, next, target)
			for _, child := range children {
				out = append(out, toString(direction)+child)
			}
		}
	}
	return out
}

func shortestLength() func(a string, b string) int {
	return func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
}

func toString(direction Vector) string {
	switch direction {
	case Up:
		return "^"
	case Right:
		return ">"
	case Down:
		return "v"
	case Left:
		return "<"
	}
	panic("invalid direction")
}
