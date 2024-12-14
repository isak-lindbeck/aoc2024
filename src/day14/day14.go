package day14

import (
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"math"
	"regexp"
	"strings"
)

var width = 101
var height = 103

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	split := strings.Split(input, "\n")

	if len(split) < 20 {
		width = 11
		height = 7
	}

	re := regexp.MustCompile("[-0-9]+")

	robots := make([]Robot, len(split))
	for i, s := range split {
		nums := re.FindAllString(s, -1)
		px := ints.Parse(nums[0])
		py := ints.Parse(nums[1])
		vx := ints.Parse(nums[2])
		vy := ints.Parse(nums[3])
		r := Robot{px, py, vx, vy}
		robots[i] = r

	}
	ans1 = calculateThreatLevelAtSecond(robots, 100)

	minThreatLevel := math.MaxInt32
	for i := range width * height {
		threatLevel := calculateThreatLevelAtSecond(robots, i)
		if threatLevel < minThreatLevel {
			minThreatLevel = threatLevel
			ans2 = i
		}
	}

	return ans1, ans2
}

func calculateThreatLevelAtSecond(robots []Robot, seconds int) int {
	quadrants := []int{0, 0, 0, 0}
	for _, r := range robots {
		r.move(seconds)
		onMiddle := r.px == width/2 || r.py == height/2
		onLeftHalf := r.px < width/2
		onUpperHalf := r.py < height/2
		quadIdx := ints.FromBool(onLeftHalf)*2 + ints.FromBool(onUpperHalf)
		quadrants[quadIdx] += ints.FromBool(!onMiddle)
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

type Robot struct {
	px, py, vx, vy int
}

func (r *Robot) move(seconds int) {
	r.px = (((r.px + r.vx*seconds) % width) + width) % width
	r.py = (((r.py + r.vy*seconds) % height) + height) % height
}
