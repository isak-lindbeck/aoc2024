package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/isak-lindbeck/aoc2024/src/day01"
	"github.com/isak-lindbeck/aoc2024/src/day02"
	"github.com/isak-lindbeck/aoc2024/src/day03"
	"github.com/isak-lindbeck/aoc2024/src/day04"
	"github.com/isak-lindbeck/aoc2024/src/day05"
	"github.com/isak-lindbeck/aoc2024/src/day06"
	"github.com/isak-lindbeck/aoc2024/src/day07"
	"github.com/isak-lindbeck/aoc2024/src/day08"
	"github.com/isak-lindbeck/aoc2024/src/day09"
	"github.com/isak-lindbeck/aoc2024/src/day10"
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"os"
	"strings"
	"time"
)

func main() {
	today := time.Now().Day()

	runDay(today)

	if false {
		for day := 1; day < min(today, 25); day++ {
			runDay(day)
		}
	}
}

func runDay(day int) (int, int) {
	inputAsString := utils.InputAsString(fmt.Sprintf("inputs/day%02d.txt", day))

	ans1 := 0
	ans2 := 0
	start := time.Now()
	switch day {
	case 1:
		ans1, ans2 = day01.Run(inputAsString)
	case 2:
		ans1, ans2 = day02.Run(inputAsString)
	case 3:
		ans1, ans2 = day03.Run(inputAsString)
	case 4:
		ans1, ans2 = day04.Run(inputAsString)
	case 5:
		ans1, ans2 = day05.Run(inputAsString)
	case 6:
		ans1, ans2 = day06.Run(inputAsString)
	case 7:
		ans1, ans2 = day07.Run(inputAsString)
	case 8:
		ans1, ans2 = day08.Run(inputAsString)
	case 9:
		ans1, ans2 = day09.Run(inputAsString)
	case 10:
		ans1, ans2 = day10.Run(inputAsString)
	}
	duration("Runtime:", start)

	checkAnswers(ans1, ans2, day)
	return ans1, ans2
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

func checkAnswers(ans1 int, ans2 int, day int) {

	input, err := os.ReadFile(fmt.Sprintf("answers/day%02d.txt", day))
	if err == nil {
		answers := string(input)
		split := strings.Split(answers, "\n")
		fmt.Printf("Answers day%02d:\n", day)
		if ints.Parse(split[0]) != ans1 {
			color.Red("%d\n", ans1)
		} else {
			color.Green("%d\n", ans1)
		}
		if ints.Parse(split[1]) != ans2 {
			color.Red("%d\n", ans2)
		} else {
			color.Green("%d\n", ans2)
		}
	} else {
		fmt.Printf("Answers day%02d:\n", day)
		fmt.Println(ans1)
		fmt.Println(ans2)
	}
	fmt.Println()

}
