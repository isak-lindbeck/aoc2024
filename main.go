package main

import (
	"fmt"
	"github.com/isak-lindbeck/aoc2024/src/day06"
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"log"
	"time"
)

func main() {
	defer duration(track("runtime"))

	p1, p2 := day06.Run(utils.InputAsString("inputs/day06.txt"))

	fmt.Println(p1)
	fmt.Println(p2)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
