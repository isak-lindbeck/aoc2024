package day11

import (
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"strconv"
	"strings"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	slice := ints.ParseSlice(strings.Fields(input))
	cache := make(map[Key]int)

	for _, stone := range slice {
		ans1 += blink(stone, 25, cache)
	}

	for _, stone := range slice {
		ans2 += blink(stone, 75, cache)
	}

	return ans1, ans2
}

func blink(stone, level int, cache map[Key]int) int {
	level--
	if level < 0 {
		return 1
	}
	key := Key{stone: stone, level: level}

	if value, exist := cache[key]; exist {
		return value
	}

	if stone == 0 {
		cache[key] = blink(1, level, cache)
		return cache[key]
	}
	stoneString := strconv.Itoa(stone)
	stoneDigits := len(stoneString)
	if stoneDigits%2 == 0 {
		pow := ints.Pow(10, stoneDigits/2)
		cache[key] = blink(stone/pow, level, cache) + blink(stone%pow, level, cache)
		return cache[key]
	}
	cache[key] = blink(stone*2024, level, cache)
	return cache[key]
}

type Key struct {
	stone, level int
}
