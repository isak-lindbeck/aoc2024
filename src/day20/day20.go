package day20

import (
	"fmt"
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
	matrix.SetAt(start.xy())('.')
	matrix.SetAt(end.xy())('.')
	outsideMap := Vector{matrix.Width, matrix.Height}
	original := Dijkstra{
		utils.NewMatrixWithDefaultValue(matrix.Width, matrix.Height, math.MaxInt),
		utils.NewMatrix[Vector](matrix.Width, matrix.Height),
		utils.NewQueue(make([]Vector, matrix.Width*matrix.Height)),
	}
	original.dist.Set(start.x, start.y, 0)
	original.queue.PushFront(start)
	Solve(original, matrix, outsideMap, outsideMap, 0)

	ans1 = doIt(matrix, original, end, 2)
	ans2 = doIt(matrix, original, end, 20)

	return ans1, ans2
}

func doIt(matrix utils.Matrix[rune], original Dijkstra, end Vector, maxCheatDistance int) int {
	sum := 0
	expectedSave := 100
	savedDistance := make(map[int]int)
	total := (matrix.Height - 2) * (matrix.Width - 2)
	count := 0
	for y := 1; y < matrix.Height-1; y++ {
		for x := 1; x < matrix.Width-1; x++ {
			cheatStart := Vector{x, y}
			if matrix.Get(cheatStart.xy()) != '.' {
				continue
			}
			for dy := -maxCheatDistance; dy <= maxCheatDistance; dy++ {
				for dx := -maxCheatDistance; dx <= maxCheatDistance; dx++ {
					cheatEnd := Vector{x + dx, y + dy}
					cheatDistance := cheatStart.distance(cheatEnd)
					if matrix.GetSafeAt(cheatEnd.xy())('X') != '.' {
						continue
					}
					if cheatDistance < 2 {
						continue
					}
					if cheatDistance > maxCheatDistance {
						continue
					}
					newState := original.Clone()
					newState.queue.PushBack(cheatStart)
					Solve(newState, matrix, cheatStart, cheatEnd, expectedSave-1)
					saved := original.dist.Get(end.xy()) - newState.dist.Get(end.xy())
					savedDistance[saved]++
					if saved >= expectedSave {
						sum++
					}
				}
			}
			count++
			if count%100 == 0 {
				fmt.Print(count)
				fmt.Print(" / ")
				fmt.Println(total)
			}
		}
	}

	return sum
}

type Dijkstra struct {
	dist  utils.Matrix[int]
	prev  utils.Matrix[Vector]
	queue utils.Queue[Vector]
}

func (d Dijkstra) Clone() Dijkstra {
	return Dijkstra{
		utils.CloneMatrix(d.dist),
		utils.CloneMatrix(d.prev),
		d.queue.Clone(),
	}
}

func Solve(state Dijkstra, matrix utils.Matrix[rune], cheatStart, cheatEnd Vector, expectedSave int) Dijkstra {

	for true {
		cur, exists := state.queue.Pop()
		if !exists {
			break
		}
		curDist := state.dist.Get(cur.xy())

		for _, direction := range directions {
			next := cur.add(direction)
			if matrix.GetSafeAt(next.xy())('#') != '#' {
				alt := curDist + 1
				if alt < state.dist.Get(next.xy())-expectedSave {
					state.dist.SetAt(next.xy())(alt)
					state.prev.SetAt(next.xy())(cur)
					state.queue.PushBack(next)
				}
			}
		}
		if cur == cheatStart {
			next := cheatEnd
			cheatDistance := cheatStart.distance(cheatEnd)
			alt := curDist + cheatDistance
			if alt < state.dist.Get(next.xy())-expectedSave {
				state.dist.SetAt(next.xy())(alt)
				state.prev.SetAt(next.xy())(cur)
				state.queue.PushBack(next)
			}
		}
	}

	return state
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
