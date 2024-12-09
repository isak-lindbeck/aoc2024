package day08

import (
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"strings"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	matrix := utils.RuneMatrix(strings.TrimSuffix(input, "\n"))

	asn1Matrix := utils.NewMatrixWithDefaultValue(matrix.Width, matrix.Height, false)
	asn2Matrix := utils.NewMatrixWithDefaultValue(matrix.Width, matrix.Height, false)

	var frequencyMap = make(map[rune][]Point)

	for x, y := range matrix.Keys() {
		r := matrix.Get(x, y)
		if r != '.' {
			frequencyMap[r] = append(frequencyMap[r], Point{x, y})
		}
	}

	for _, points := range frequencyMap {
		for _, pointA := range points {
			for _, pointB := range points {
				if pointA == pointB {
					continue
				}

				dx := pointB.x - pointA.x
				dy := pointB.y - pointA.y
				asn1Matrix.SetSafe(pointB.x+dx, pointB.y+dy, true)

				point := Point{pointB.x, pointB.y}
				for asn2Matrix.SetSafe(point.x, point.y, true) {
					point.x = point.x + dx
					point.y = point.y + dy
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

type Point struct {
	x, y int
}
