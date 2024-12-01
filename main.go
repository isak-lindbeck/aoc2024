package main

import (
	"fmt"
	"github.com/isak-lindbeck/aoc2024/src/day01"
	"github.com/isak-lindbeck/aoc2024/src/utils"
)

func main() {
	p1, p2 := day01.Run(utils.InputAsString("inputs/day01.txt"))

	fmt.Println(p1)
	fmt.Println(p2)
}
