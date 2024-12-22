package day06

import (
	. "github.com/isak-lindbeck/aoc2024/src/utils"
)

var directions = []Vector{
	Up,
	Right,
	Down,
	Left,
}

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	matrix := RuneMatrix(input)
	currentDirection := 0
	start := NewVector(matrix.GetCoordinates('^'))
	var visited = make(map[Vector]bool)
	currentPos := start
	for matrix.GetSafeAt(currentPos.XY())('X') != 'X' {
		visited[currentPos] = true
		nextPos := currentPos.Add(directions[currentDirection].XY())
		if matrix.GetSafeAt(nextPos.XY())('X') == '#' {
			currentDirection = (currentDirection + 1) % 4
		} else {
			currentPos = nextPos
		}

	}
	ans1 = len(visited)

	regularPath := visited
	delete(regularPath, start)

	c := make(chan int)
	for block := range regularPath {
		go calculateBlockedPosition(matrix, start, block, c)
	}
	for range regularPath {
		ans2 += <-c
	}

	return ans1, ans2
}

func calculateBlockedPosition(matrix Matrix[rune], start Vector, block Vector, c chan int) {
	sum := 0
	visitedWithDirection := NewMatrixWithDefaultValue(matrix.Width, matrix.Height, uint8(0))

	currentDirection := 0
	currentPos := start
	for matrix.GetSafeAt(currentPos.XY())('X') != 'X' {
		visitedBitFlags := visitedWithDirection.Get(currentPos.X, currentPos.Y)
		dirBitFlag := uint8(1) << currentDirection
		if (visitedBitFlags & dirBitFlag) > 0 {
			sum++
			break
		} else {
			visitedWithDirection.SetAt(currentPos.X, currentPos.Y)(visitedBitFlags | dirBitFlag)
		}
		nextPos := currentPos.Add(directions[currentDirection].XY())
		if matrix.GetSafeAt(nextPos.XY())('X') == '#' || nextPos == block {
			currentDirection = (currentDirection + 1) % 4
		} else {
			currentPos = nextPos
		}
	}
	c <- sum
}
