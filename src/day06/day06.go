package day06

import (
	. "github.com/isak-lindbeck/aoc2024/src/utils"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	directions := []Vector{
		{0, -1}, // ↑
		{1, 0},  // →
		{0, 1},  // ↓
		{-1, 0}, // ←
	}
	matrix := RuneMatrix(input)
	currentDirection := 0
	startX, startY := matrix.GetCoordinates('^')
	var visited = make(map[Vector]bool)
	currentPos := Vector{X: startX, Y: startY}
	for matrix.GetSafe(currentPos.X, currentPos.Y, 'X') != 'X' {
		visited[currentPos] = true
		next := matrix.GetSafe(currentPos.X+directions[currentDirection].X, currentPos.Y+directions[currentDirection].Y, 'X')
		if next == '#' {
			currentDirection = (currentDirection + 1) % 4
		} else {
			currentPos = currentPos.Add(directions[currentDirection].XY())
		}

	}
	ans1 = len(visited)

	regularPath := visited
	delete(regularPath, Vector{startX, startY})

	for blockPos := range regularPath {
		matrix.Set(blockPos.X, blockPos.Y, '#')

		currentDirection = 0
		visitedWithDirection := NewMatrixWithDefaultValue(matrix.Width, matrix.Height, 0)

		currentPos = Vector{X: startX, Y: startY}
		for matrix.GetSafe(currentPos.X, currentPos.Y, 'X') != 'X' {
			visitedBitFlags := visitedWithDirection.Get(currentPos.X, currentPos.Y)
			if visitedBitFlags&(1<<currentDirection) > 0 {
				ans2++
				break
			} else {
				visitedWithDirection.Set(currentPos.X, currentPos.Y, visitedBitFlags|(1<<currentDirection))
			}

			next := matrix.GetSafe(currentPos.X+directions[currentDirection].X, currentPos.Y+directions[currentDirection].Y, 'X')
			if next == '#' {
				currentDirection = (currentDirection + 1) % 4
			} else {
				currentPos = currentPos.Add(directions[currentDirection].XY())
			}
		}

		matrix.Set(blockPos.X, blockPos.Y, '.')
	}

	return ans1, ans2
}
