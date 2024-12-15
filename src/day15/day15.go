package day15

import (
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"strings"
)

var up = Instruction{0, -1}
var down = Instruction{0, 1}
var left = Instruction{-1, 0}
var right = Instruction{1, 0}

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	split := strings.Split(input, "\n\n")

	instructionString := strings.ReplaceAll(split[1], "\n", "")
	instructions := make([]Instruction, len(instructionString))
	for i, r := range instructionString {
		instructions[i] = toInstruction(r)
	}

	matrix := utils.RuneMatrix(split[0])
	atX, atY := matrix.GetCoordinates('@')
	matrix.Set(atX, atY, '.')

	for _, instr := range instructions {
		newMatrix := utils.CloneMatrix(matrix)
		if push(&newMatrix, instr, atX+instr.x, atY+instr.y) {
			atX += instr.x
			atY += instr.y
			matrix = newMatrix
		}
	}
	ans1 = sumGps(matrix)

	m := strings.ReplaceAll(split[0], ".", "..")
	m = strings.ReplaceAll(m, "#", "##")
	m = strings.ReplaceAll(m, "O", "[]")
	m = strings.ReplaceAll(m, "@", "@.")
	matrix = utils.RuneMatrix(m)
	atX, atY = matrix.GetCoordinates('@')
	matrix.Set(atX, atY, '.')
	for _, instr := range instructions {
		newMatrix := utils.CloneMatrix(matrix)
		if push(&newMatrix, instr, atX+instr.x, atY+instr.y) {
			atX += instr.x
			atY += instr.y
			matrix = newMatrix
		}
	}
	ans2 = sumGps(matrix)

	return ans1, ans2
}

func push(matrix *utils.Matrix[rune], dir Instruction, x, y int) bool {
	tile := matrix.Get(x, y)
	if tile == '#' {
		return false
	}
	if tile == '.' {
		return true
	}
	isVertical := dir == up || dir == down
	nx := x + dir.x
	ny := y + dir.y
	if isVertical && tile == '[' {
		if push(matrix, dir, nx+1, ny) {
			matrix.Set(nx+1, ny, ']')
			matrix.Set(x+1, y, '.')
		} else {
			return false
		}
	}
	if isVertical && tile == ']' {
		if push(matrix, dir, nx-1, ny) {
			matrix.Set(nx-1, ny, '[')
			matrix.Set(x-1, y, '.')
		} else {
			return false
		}
	}

	if push(matrix, dir, nx, ny) {
		matrix.Set(nx, ny, tile)
		matrix.Set(x, y, '.')
		return true
	} else {
		return false
	}
}

func sumGps(matrix utils.Matrix[rune]) int {
	sum := 0
	for x, y := range matrix.Keys() {
		if matrix.Get(x, y) == 'O' || matrix.Get(x, y) == '[' {
			sum += 100*y + x
		}
	}
	return sum
}

func toInstruction(r int32) Instruction {
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

type Instruction struct{ x, y int }
