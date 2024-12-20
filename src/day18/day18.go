package day18

import (
	"fmt"
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"math"
	"regexp"
)

var side = 71

var up = Vector{0, -1}
var right = Vector{1, 0}
var down = Vector{0, 1}
var left = Vector{-1, 0}

var directions = []Vector{up, right, down, left}

func Run(input string) (int, string) {
	ans1 := 0
	ans2 := ""

	re := regexp.MustCompile("[-0-9]+")
	allNumbers := re.FindAllString(input, -1)

	partOneNumBlocks := 1024
	if len(allNumbers) < 100 {
		side = 7
		partOneNumBlocks = 12
	}
	from := Vector{0, 0}
	to := Vector{side - 1, side - 1}

	blocks := make([]Vector, len(allNumbers)/2)
	for i := 0; i < len(allNumbers); i += 2 {
		x := ints.Parse(allNumbers[i])
		y := ints.Parse(allNumbers[i+1])
		blocks[i/2] = Vector{x, y}
	}

	matrix := utils.NewMatrixWithDefaultValue(side, side, true)
	part1Matrix := addBlocks(matrix, blocks, 0, partOneNumBlocks)

	ans1 = Dijkstra(part1Matrix, from, to)

	matrix = utils.NewMatrixWithDefaultValue(side, side, true)
	low := 0
	high := len(blocks)
	for low <= high {
		mid := (low + high) / 2
		newMatrix := addBlocks(utils.CloneMatrix(matrix), blocks, 0, mid)
		distance1 := Dijkstra(newMatrix, from, to)
		newMatrix = addBlocks(newMatrix, blocks, mid, mid+1)
		distance2 := Dijkstra(newMatrix, from, to)
		maxInt := math.MaxInt

		if distance1 < math.MaxInt && distance2 == math.MaxInt {
			break
		}

		if distance1 < maxInt {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	x, y := blocks[low].xy()
	ans2 = fmt.Sprintf("%d,%d", x, y)

	return ans1, ans2
}

func addBlocks(matrix utils.Matrix[bool], blocks []Vector, fromIdx, toIdx int) utils.Matrix[bool] {
	for i := fromIdx; i < toIdx; i++ {
		x, y := blocks[i].xy()
		matrix.Set(x, y, false)
	}
	return matrix
}

func Dijkstra(matrix utils.Matrix[bool], from, to Vector) int {
	dist := utils.NewMatrixWithDefaultValue(side, side, math.MaxInt)
	var prev = utils.NewMatrix[Vector](side, side)
	dist.Set(from.x, from.y, 0)

	queue := utils.NewQueue(make([]Vector, side))
	queue.PushFront(from)

	for true {
		cur, exists := queue.Pop()
		if !exists {
			break
		}
		curDist := dist.Get(cur.xy())

		for _, direction := range directions {
			next := cur.add(direction)
			x, y := next.xy()
			if matrix.GetSafe(x, y, false) {
				alt := curDist + 1
				if alt < dist.Get(x, y) {
					dist.Set(x, y, alt)
					prev.Set(x, y, cur)
					queue.PushBack(next)
				}
			}
		}
	}

	return dist.Get(to.xy())
}

type Vector struct{ x, y int }

func (v *Vector) xy() (int, int) {
	return v.x, v.y
}

func (v1 *Vector) add(v2 Vector) Vector {
	return Vector{v1.x + v2.x, v1.y + v2.y}
}
