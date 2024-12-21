package day08

import (
	. "github.com/isak-lindbeck/aoc2024/src/utils"
	"strings"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	matrix := RuneMatrix(strings.TrimSuffix(input, "\n"))

	asn1Matrix := NewMatrixWithDefaultValue(matrix.Width, matrix.Height, false)
	asn2Matrix := NewMatrixWithDefaultValue(matrix.Width, matrix.Height, false)

	var frequencyMap = make(map[rune][]Vector)

	for x, y := range matrix.Keys() {
		r := matrix.Get(x, y)
		if r != '.' {
			frequencyMap[r] = append(frequencyMap[r], Vector{x, y})
		}
	}

	for _, points := range frequencyMap {
		for _, pointA := range points {
			for _, pointB := range points {
				if pointA == pointB {
					continue
				}

				dx := pointB.X - pointA.X
				dy := pointB.Y - pointA.Y
				asn1Matrix.SetSafe(pointB.X+dx, pointB.Y+dy, true)

				point := Vector{pointB.X, pointB.Y}
				for asn2Matrix.SetSafe(point.X, point.Y, true) {
					point.X = point.X + dx
					point.Y = point.Y + dy
				}
			}
		}
	}

	for x, y := range matrix.Keys() {
		if asn1Matrix.Get(x, y) {
			ans1++
		}
		if asn2Matrix.Get(x, y) {
			ans2++
		}
	}

	return ans1, ans2
}
