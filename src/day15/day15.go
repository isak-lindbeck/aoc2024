package day15

import (
	. "github.com/isak-lindbeck/aoc2024/src/utils"
	"strings"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	split := strings.Split(input, "\n\n")

	instructionString := strings.ReplaceAll(split[1], "\n", "")
	instructions := make([]Vector, len(instructionString))
	for i, r := range instructionString {
		instructions[i] = toInstruction(r)
	}

	ans1 = calculate(split[0], instructions)

	m := strings.ReplaceAll(split[0], ".", "..")
	m = strings.ReplaceAll(m, "#", "##")
	m = strings.ReplaceAll(m, "O", "[]")
	m = strings.ReplaceAll(m, "@", "@.")
	ans2 = calculate(m, instructions)

	return ans1, ans2
}

func calculate(input string, instructions []Vector) int {
	matrix := RuneMatrix(input)
	atX, atY := matrix.GetCoordinates('@')
	matrix.Set(atX, atY, '.')
	for _, instr := range instructions {
		if canPush(&matrix, instr, atX+instr.X, atY+instr.Y) {
			push(&matrix, instr, atX+instr.X, atY+instr.Y)
			atX += instr.X
			atY += instr.Y
		}
	}
	sum := 0
	for x, y := range matrix.Keys() {
		if matrix.Get(x, y) == 'O' || matrix.Get(x, y) == '[' {
			sum += 100*y + x
		}
	}
	return sum
}

func canPush(matrix *Matrix[rune], dir Vector, x, y int) bool {
	tile := matrix.Get(x, y)
	if tile == '#' {
		return false
	}
	if tile == '.' {
		return true
	}
	nextX := x + dir.X
	nextY := y + dir.Y
	if dir.X == 0 && tile == '[' {
		if !canPush(matrix, dir, nextX+1, nextY) {
			return false
		}
	}
	if dir.X == 0 && tile == ']' {
		if !canPush(matrix, dir, nextX-1, nextY) {
			return false
		}
	}
	return canPush(matrix, dir, nextX, nextY)
}

func push(matrix *Matrix[rune], dir Vector, x, y int) {
	tile := matrix.Get(x, y)
	if tile == '#' || tile == '.' {
		return
	}
	nextX := x + dir.X
	nextY := y + dir.Y
	if dir.X == 0 && tile == '[' {
		push(matrix, dir, nextX+1, nextY)
		matrix.Set(nextX+1, nextY, ']')
		matrix.Set(x+1, y, '.')

	}
	if dir.X == 0 && tile == ']' {
		push(matrix, dir, nextX-1, nextY)
		matrix.Set(nextX-1, nextY, '[')
		matrix.Set(x-1, y, '.')
	}

	push(matrix, dir, nextX, nextY)
	matrix.Set(nextX, nextY, tile)
	matrix.Set(x, y, '.')
}

func toInstruction(r rune) Vector {
	switch r {
	case '^':
		return Up
	case 'v':
		return Down
	case '<':
		return Left
	case '>':
		return Right
	default:
		panic("Invalid character in input")
	}
}
