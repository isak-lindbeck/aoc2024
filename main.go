package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/isak-lindbeck/aoc2024/src/day07"
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	defer duration(track("runtime"))

	day := "07"

	ans1, ans2 := day07.Run(utils.InputAsString(fmt.Sprintf("inputs/day%s.txt", day)))

	checkAnswers(ans1, ans2, day)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func checkAnswers(ans1 int, ans2 int, day string) {
	fmt.Println()
	fmt.Println("Answers:")
	fmt.Println(ans1)
	fmt.Println(ans2)
	fmt.Println()
	input, err := os.ReadFile(fmt.Sprintf("answers/day%s.txt", day))
	if err == nil {
		answers := string(input)
		split := strings.Split(answers, "\n")
		if ints.Parse(split[0]) != ans1 {
			color.Red("Answer one '%s' is wrong! Expected '%d'\n", split[0], ans1)
		}
		if ints.Parse(split[1]) != ans2 {
			color.Red("Answer two '%s' is wrong! Expected '%d'\n", split[1], ans2)
		}
		fmt.Println()
	}
}
