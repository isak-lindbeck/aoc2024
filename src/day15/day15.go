package day15

import (
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"strings"
)

var up = Vector{0, -1}
var down = Vector{0, 1}
var left = Vector{-1, 0}
var right = Vector{1, 0}

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
	matrix := utils.RuneMatrix(input)
	atX, atY := matrix.GetCoordinates('@')
	matrix.Set(atX, atY, '.')
	for _, instr := range instructions {
		if canPush(&matrix, instr, atX+instr.x, atY+instr.y) {
			push(&matrix, instr, atX+instr.x, atY+instr.y)
			atX += instr.x
			atY += instr.y
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

func canPush(matrix *utils.Matrix[rune], dir Vector, x, y int) bool {
	tile := matrix.Get(x, y)
	if tile == '#' {
		return false
	}
	if tile == '.' {
		return true
	}
	nextX := x + dir.x
	nextY := y + dir.y
	if dir.x == 0 && tile == '[' {
		if !canPush(matrix, dir, nextX+1, nextY) {
			return false
		}
	}
	if dir.x == 0 && tile == ']' {
		if !canPush(matrix, dir, nextX-1, nextY) {
			return false
		}
	}
	return canPush(matrix, dir, nextX, nextY)
}

func push(matrix *utils.Matrix[rune], dir Vector, x, y int) {
	tile := matrix.Get(x, y)
	if tile == '#' || tile == '.' {
		return
	}
	nextX := x + dir.x
	nextY := y + dir.y
	if dir.x == 0 && tile == '[' {
		push(matrix, dir, nextX+1, nextY)
		matrix.Set(nextX+1, nextY, ']')
		matrix.Set(x+1, y, '.')

	}
	if dir.x == 0 && tile == ']' {
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
		return up
	case 'v':
		return down
	case '<':
		return left
	case '>':
		return right
	default:
		panic("Invalid character in input")
	}
}

type Vector struct{ x, y int }

func (v1 *Vector) move(v2 Vector) {
	v1.x += v2.x
	v1.y += v2.y
}
