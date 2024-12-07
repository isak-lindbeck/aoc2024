package day06

import (
	"github.com/isak-lindbeck/aoc2024/src/utils"
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
	matrix := utils.RuneMatrix(input)
	currentDirection := 0
	startX, startY := matrix.GetCoordinates('^')
	var visited = make(map[Vector]bool)
	currentPos := Vector{x: startX, y: startY}
	for matrix.GetSafe(currentPos.x, currentPos.y, 'X') != 'X' {
		visited[currentPos] = true
		next := matrix.GetSafe(currentPos.x+directions[currentDirection].x, currentPos.y+directions[currentDirection].y, 'X')
		if next == '#' {
			currentDirection = (currentDirection + 1) % 4
		} else {
			currentPos.move(directions[currentDirection])
		}

	}
	ans1 = len(visited)

	regularPath := visited
	delete(regularPath, Vector{startX, startY})

	for blockPos := range regularPath {
		matrix.Set(blockPos.x, blockPos.y, '#')

		currentDirection = 0
		visitedWithDirection := utils.NewMatrixWithDefaultValue(matrix.Width, matrix.Height, 0)

		currentPos = Vector{x: startX, y: startY}
		for matrix.GetSafe(currentPos.x, currentPos.y, 'X') != 'X' {
			visitedBitFlags := visitedWithDirection.Get(currentPos.x, currentPos.y)
			if visitedBitFlags&(1<<currentDirection) > 0 {
				ans2++
				break
			} else {
				visitedWithDirection.Set(currentPos.x, currentPos.y, visitedBitFlags|(1<<currentDirection))
			}

			next := matrix.GetSafe(currentPos.x+directions[currentDirection].x, currentPos.y+directions[currentDirection].y, 'X')
			if next == '#' {
				currentDirection = (currentDirection + 1) % 4
			} else {
				currentPos.move(directions[currentDirection])
			}
		}

		matrix.Set(blockPos.x, blockPos.y, '.')
	}

	return ans1, ans2
}

type Vector struct{ x, y int }

func (v1 *Vector) move(v2 Vector) {
	v1.x += v2.x
	v1.y += v2.y
}
