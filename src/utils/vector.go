package utils

import "github.com/isak-lindbeck/aoc2024/src/ints"

var Up = Vector{Y: -1}
var Right = Vector{X: 1}
var Down = Vector{Y: 1}
var Left = Vector{X: -1}

var North = Up
var East = Right
var South = Down
var West = Left

var NorthEast = North.Add(East.XY())
var SouthEast = South.Add(East.XY())
var SouthWest = South.Add(West.XY())
var NorthWest = North.Add(West.XY())

type Vector struct{ X, Y int }

func NewVector(x, y int) Vector {
	return Vector{x, y}
}

func (v *Vector) XY() (int, int) {
	return v.X, v.Y
}

func (v1 *Vector) Add(x, y int) Vector {
	return Vector{v1.X + x, v1.Y + y}
}

func (v1 *Vector) Subtract(x, y int) Vector {
	return Vector{v1.X - x, v1.Y - y}
}

func (v Vector) Mul(a int) Vector {
	return Vector{v.X * a, v.Y * a}
}

func (v1 *Vector) Distance(v2 Vector) int {
	return ints.Abs(v1.X-v2.X) + ints.Abs(v1.Y-v2.Y)
}
