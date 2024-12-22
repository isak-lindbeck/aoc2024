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
	"github.com/isak-lindbeck/aoc2024/src/day11"
	"github.com/isak-lindbeck/aoc2024/src/day12"
	"github.com/isak-lindbeck/aoc2024/src/day13"
	"github.com/isak-lindbeck/aoc2024/src/day14"
	"github.com/isak-lindbeck/aoc2024/src/day15"
	"github.com/isak-lindbeck/aoc2024/src/day16"
	"github.com/isak-lindbeck/aoc2024/src/day17"
	"github.com/isak-lindbeck/aoc2024/src/day18"
	"github.com/isak-lindbeck/aoc2024/src/day19"
	"github.com/isak-lindbeck/aoc2024/src/day20"
	"github.com/isak-lindbeck/aoc2024/src/day21"
	"github.com/isak-lindbeck/aoc2024/src/day22"
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"os"
	"strconv"
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

func runDay(day int) (string, string) {
	inputAsString := utils.InputAsString(fmt.Sprintf("inputs/day%02d.txt", day))
	inputAsString = strings.TrimSuffix(inputAsString, "\n")

	ans1 := "0"
	ans2 := "0"
	start := time.Now()
	switch day {
	case 1:
		ans1, ans2 = asStringAns(day01.Run(inputAsString))
	case 2:
		ans1, ans2 = asStringAns(day02.Run(inputAsString))
	case 3:
		ans1, ans2 = asStringAns(day03.Run(inputAsString))
	case 4:
		ans1, ans2 = asStringAns(day04.Run(inputAsString))
	case 5:
		ans1, ans2 = asStringAns(day05.Run(inputAsString))
	case 6:
		ans1, ans2 = asStringAns(day06.Run(inputAsString))
	case 7:
		ans1, ans2 = asStringAns(day07.Run(inputAsString))
	case 8:
		ans1, ans2 = asStringAns(day08.Run(inputAsString))
	case 9:
		ans1, ans2 = asStringAns(day09.Run(inputAsString))
	case 10:
		ans1, ans2 = asStringAns(day10.Run(inputAsString))
	case 11:
		ans1, ans2 = asStringAns(day11.Run(inputAsString))
	case 12:
		ans1, ans2 = asStringAns(day12.Run(inputAsString))
	case 13:
		ans1, ans2 = asStringAns(day13.Run(inputAsString))
	case 14:
		ans1, ans2 = asStringAns(day14.Run(inputAsString))
	case 15:
		ans1, ans2 = asStringAns(day15.Run(inputAsString))
	case 16:
		ans1, ans2 = asStringAns(day16.Run(inputAsString))
	case 17:
		ans1, ans2 = asStringAns2(day17.Run(inputAsString))
	case 18:
		ans1, ans2 = asStringAns1(day18.Run(inputAsString))
	case 19:
		ans1, ans2 = asStringAns(day19.Run(inputAsString))
	case 20:
		ans1, ans2 = asStringAns(day20.Run(inputAsString))
	case 21:
		ans1, ans2 = asStringAns(day21.Run(inputAsString))
	case 22:
		ans1, ans2 = asStringAns(day22.Run(inputAsString))
	default:
		color.Red("Day %d not implemented!\n", day)
	}
	duration("Runtime:", start)

	checkAnswers(ans1, ans2, day)
	return ans1, ans2
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

func checkAnswers(ans1, ans2 string, day int) {

	input, err := os.ReadFile(fmt.Sprintf("answers/day%02d.txt", day))
	if err == nil {
		answers := string(input)
		split := strings.Split(answers, "\n")
		fmt.Printf("Answers day%02d:\n", day)
		if split[0] != ans1 {
			color.Red("%s\n", ans1)
		} else {
			color.Green("%s\n", ans1)
		}
		if split[1] != ans2 {
			color.Red("%s\n", ans2)
		} else {
			color.Green("%s\n", ans2)
		}
	} else {
		fmt.Printf("Answers day%02d:\n", day)
		fmt.Println(ans1)
		fmt.Println(ans2)
	}
	fmt.Println()

}

func asStringAns(a int, b int) (string, string) {
	return strconv.Itoa(a), strconv.Itoa(b)
}

func asStringAns1(a int, b string) (string, string) {
	return strconv.Itoa(a), b
}

func asStringAns2(a string, b int) (string, string) {
	return a, strconv.Itoa(b)
}
