package day20

import (
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"math"
)

var up = Vector{0, -1}
var right = Vector{1, 0}
var down = Vector{0, 1}
var left = Vector{-1, 0}

var directions = []Vector{up, right, down, left}

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	matrix := utils.RuneMatrix(input)
	start := NewVector(matrix.GetCoordinates('S'))
	end := NewVector(matrix.GetCoordinates('E'))

	distFromStart := SolveDijkstra(matrix, start)
	distFromEnd := SolveDijkstra(matrix, end)

	ans1 = calculateCheatSuccesses(matrix, distFromStart, distFromEnd, end, 2, 100)
	ans2 = calculateCheatSuccesses(matrix, distFromStart, distFromEnd, end, 20, 100)

	return ans1, ans2
}

func calculateCheatSuccesses(matrix utils.Matrix[rune], distFromStart utils.Matrix[int], distF utils.Matrix[int], end Vector, maxCheatDistance, expectedSavedDistance int) int {
	sum := 0
	for y := 1; y < matrix.Height-1; y++ {
		for x := 1; x < matrix.Width-1; x++ {
			cheatStart := Vector{x, y}
			if matrix.Get(cheatStart.xy()) == '#' {
				continue
			}
			for dy := -maxCheatDistance; dy <= maxCheatDistance; dy++ {
				for dx := -maxCheatDistance; dx <= maxCheatDistance; dx++ {
					cheatEnd := Vector{x + dx, y + dy}
					if matrix.GetSafeAt(cheatEnd.xy())('#') == '#' {
						continue
					}
					cheatDistance := cheatStart.distance(cheatEnd)
					if cheatDistance < 2 || cheatDistance > maxCheatDistance {
						continue
					}
					oldDist := distFromStart.Get(end.xy())
					newDist := distFromStart.Get(cheatStart.xy()) + cheatDistance + distF.Get(cheatEnd.xy())
					saved := oldDist - newDist
					if saved >= expectedSavedDistance {
						sum++
					}
				}
			}
		}
	}

	return sum
}

func SolveDijkstra(matrix utils.Matrix[rune], start Vector) utils.Matrix[int] {
	dist := utils.NewMatrixWithDefaultValue(matrix.Width, matrix.Height, math.MaxInt)
	prev := utils.NewMatrix[Vector](matrix.Width, matrix.Height)
	queue := utils.NewQueue(make([]Vector, matrix.Width*matrix.Height))

	dist.Set(start.x, start.y, 0)
	queue.PushFront(start)

	for true {
		cur, exists := queue.Pop()
		if !exists {
			break
		}
		curDist := dist.Get(cur.xy())

		for _, direction := range directions {
			next := cur.add(direction)
			if matrix.GetSafeAt(next.xy())('#') != '#' {
				alt := curDist + 1
				if alt < dist.Get(next.xy()) {
					dist.SetAt(next.xy())(alt)
					prev.SetAt(next.xy())(cur)
					queue.PushBack(next)
				}
			}
		}
	}

	return dist
}

func NewVector(x, y int) Vector {
	return Vector{x, y}
}

type Vector struct{ x, y int }

func (v *Vector) xy() (int, int) {
	return v.x, v.y
}

func (v1 *Vector) add(v2 Vector) Vector {
	return Vector{v1.x + v2.x, v1.y + v2.y}
}

func (v1 *Vector) distance(v2 Vector) int {
	return ints.Abs(v1.x-v2.x) + ints.Abs(v1.y-v2.y)
}
