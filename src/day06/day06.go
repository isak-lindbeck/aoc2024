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
		vwdMap := make(map[VisitedWithDirection]bool)

		currentPos = Vector{x: startX, y: startY}
		works := false
		for matrix.GetSafe(currentPos.x, currentPos.y, 'X') != 'X' {
			vwd := VisitedWithDirection{currentPos, directions[currentDirection]}
			if vwdMap[vwd] {
				works = true
				break
			} else {
				vwdMap[vwd] = true
			}

			next := matrix.GetSafe(currentPos.x+directions[currentDirection].x, currentPos.y+directions[currentDirection].y, 'X')
			if next == '#' {
				currentDirection = (currentDirection + 1) % 4
			} else {
				currentPos.move(directions[currentDirection])
			}
		}

		if works {
			ans2++
		}

		matrix.Set(blockPos.x, blockPos.y, '.')
	}

	return ans1, ans2
}

type Vector struct{ x, y int }

type VisitedWithDirection struct{ v, d Vector }

func (v1 *Vector) move(v2 Vector) {
	v1.x += v2.x
	v1.y += v2.y
}
