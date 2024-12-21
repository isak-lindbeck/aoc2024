package day20

import (
	. "github.com/isak-lindbeck/aoc2024/src/utils"
	"math"
)

var directions = []Vector{Up, Right, Down, Left}

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	matrix := RuneMatrix(input)
	start := NewVector(matrix.GetCoordinates('S'))
	end := NewVector(matrix.GetCoordinates('E'))

	distFromStart := SolveDijkstra(matrix, start)
	distFromEnd := SolveDijkstra(matrix, end)

	ans1 = calculateCheatSuccesses(matrix, distFromStart, distFromEnd, end, 2, 100)
	ans2 = calculateCheatSuccesses(matrix, distFromStart, distFromEnd, end, 20, 100)

	return ans1, ans2
}

func calculateCheatSuccesses(matrix Matrix[rune], distFromStart Matrix[int], distF Matrix[int], end Vector, maxCheatDistance, expectedSavedDistance int) int {
	sum := 0
	for y := 1; y < matrix.Height-1; y++ {
		for x := 1; x < matrix.Width-1; x++ {
			cheatStart := Vector{x, y}
			if matrix.Get(cheatStart.XY()) == '#' {
				continue
			}
			for dy := -maxCheatDistance; dy <= maxCheatDistance; dy++ {
				for dx := -maxCheatDistance; dx <= maxCheatDistance; dx++ {
					cheatEnd := Vector{x + dx, y + dy}
					if matrix.GetSafeAt(cheatEnd.XY())('#') == '#' {
						continue
					}
					cheatDistance := cheatStart.Distance(cheatEnd)
					if cheatDistance < 2 || cheatDistance > maxCheatDistance {
						continue
					}
					oldDist := distFromStart.Get(end.XY())
					newDist := distFromStart.Get(cheatStart.XY()) + cheatDistance + distF.Get(cheatEnd.XY())
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

func SolveDijkstra(matrix Matrix[rune], start Vector) Matrix[int] {
	dist := NewMatrixWithDefaultValue(matrix.Width, matrix.Height, math.MaxInt)
	prev := NewMatrix[Vector](matrix.Width, matrix.Height)
	queue := NewQueue(make([]Vector, matrix.Width*matrix.Height))

	dist.Set(start.X, start.Y, 0)
	queue.PushFront(start)

	for true {
		cur, exists := queue.Pop()
		if !exists {
			break
		}
		curDist := dist.Get(cur.XY())

		for _, direction := range directions {
			next := cur.Add(direction.XY())
			if matrix.GetSafeAt(next.XY())('#') != '#' {
				alt := curDist + 1
				if alt < dist.Get(next.XY()) {
					dist.SetAt(next.XY())(alt)
					prev.SetAt(next.XY())(cur)
					queue.PushBack(next)
				}
			}
		}
	}

	return dist
}
