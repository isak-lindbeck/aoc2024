package day13

import (
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"regexp"
	"strings"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	split := strings.Split(input, "\n\n")

	re := regexp.MustCompile("[0-9]+")

	for _, s := range split {
		nums := re.FindAllString(s, -1)
		adx := ints.Parse(nums[0])
		ady := ints.Parse(nums[1])
		bdx := ints.Parse(nums[2])
		bdy := ints.Parse(nums[3])
		targetX := ints.Parse(nums[4])
		targetY := ints.Parse(nums[5])

		ans1 += calculateCost(bdy, targetX, bdx, targetY, adx, ady)
		targetX = 10000000000000 + targetX
		targetY = 10000000000000 + targetY
		ans2 += calculateCost(bdy, targetX, bdx, targetY, adx, ady)
	}

	return ans1, ans2
}

func calculateCost(bdy int, targetX int, bdx int, targetY int, adx int, ady int) int {
	a := (bdy*targetX - bdx*targetY) / (adx*bdy - bdx*ady)
	b := (adx*targetY - ady*targetX) / (adx*bdy - bdx*ady)

	isValid := (a*adx+b*bdx == targetX) && (a*ady+b*bdy == targetY)

	if isValid {
		return a*3 + b
	}
	return 0
}
