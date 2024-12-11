package day10

import (
	"github.com/isak-lindbeck/aoc2024/src/utils"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	matrix := utils.DigitMatrix(input)
	for x, y := range matrix.Keys() {
		start := matrix.Get(x, y)
		if start != 0 {
			continue
		}
		sum1, sum2 := doTheWork(matrix, x, y)

		ans1 += sum1
		ans2 += sum2
	}

	return ans1, ans2
}

func doTheWork(matrix utils.Matrix[int], x int, y int) (int, int) {
	sum1 := 0
	sum2 := 0
	visited := utils.NewMatrixWithDefaultValue(matrix.Width, matrix.Height, false)
	sum2 += travel(matrix, &visited, x, y, 0)

	for y := 0; y < visited.Height; y++ {
		for x := 0; x < visited.Width; x++ {
			if visited.Get(x, y) {
				sum1++
			}
		}
	}
	return sum1, sum2
}

func travel(m utils.Matrix[int], visited *utils.Matrix[bool], x, y, level int) int {
	value := m.GetSafe(x, y, -1)
	if value != level {
		return 0
	}
	if value == 9 {
		visited.Set(x, y, true)
		return 1
	}

	return travel(m, visited, x+1, y, level+1) +
		travel(m, visited, x-1, y, level+1) +
		travel(m, visited, x, y+1, level+1) +
		travel(m, visited, x, y-1, level+1)

}
