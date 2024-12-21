package utils

import "github.com/isak-lindbeck/aoc2024/src/ints"

type Vector struct{ X, Y int }

func NewVector(x, y int) Vector {
	return Vector{x, y}
}

func (v *Vector) XY() (int, int) {
	return v.X, v.Y
}

func (v1 *Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 *Vector) Subtract(v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 *Vector) Distance(v2 Vector) int {
	return ints.Abs(v1.X-v2.X) + ints.Abs(v1.Y-v2.Y)
}
