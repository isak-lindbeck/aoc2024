package main

import (
	"fmt"
	"github.com/isak-lindbeck/aoc2024/src/day05"
	"github.com/isak-lindbeck/aoc2024/src/utils"
)

func main() {
	p1, p2 := day05.Run(utils.InputAsString("inputs/day05.txt"))

	fmt.Println(p1)
	fmt.Println(p2)
}
