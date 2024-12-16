package day16

import (
	"github.com/isak-lindbeck/aoc2024/src/utils"
	"math"
	"slices"
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
	fromX, fromY := matrix.GetCoordinates('S')
	toX, toY := matrix.GetCoordinates('E')
	matrix.Set(fromX, fromY, '.')
	matrix.Set(toX, toY, '.')
	ans1, ans2 = Djikstra(matrix, NavPos{Vector{fromX, fromY}, 1}, Vector{toX, toY})

	return ans1, ans2
}

func Djikstra(matrix utils.Matrix[rune], from NavPos, to Vector) (int, int) {
	size := matrix.Width * matrix.Height * 4
	dist := slices.Repeat([]int{math.MaxInt}, size)
	prev := slices.Repeat([]int{-1}, size)

	dist[from.getIndex(&matrix)] = 0

	q := make([]NavPos, 1, size)
	q[0] = from
	cur := q[len(q)-1]
	for len(q) > 0 {
		cur = q[len(q)-1]
		q = q[:len(q)-1]

		next := NavPos{cur.coordinate.add(directions[cur.direction]), cur.direction}
		if matrix.Get(next.coordinate.x, next.coordinate.y) == '.' {
			alt := dist[cur.getIndex(&matrix)] + 1
			nextIdx := next.getIndex(&matrix)
			distNext := dist[nextIdx]
			if alt <= distNext {
				dist[nextIdx] = alt
				if alt < distNext {
					prev[nextIdx] = 0
				}
				prev[nextIdx] = prev[nextIdx] | 1
				q = append(q, next)
			}
		}
		next = NavPos{cur.coordinate, (cur.direction + 3) % 4}
		alt := dist[cur.getIndex(&matrix)] + 1000
		nextIdx := next.getIndex(&matrix)
		distNext := dist[nextIdx]
		if alt <= distNext {
			dist[nextIdx] = alt
			if alt < distNext {
				prev[nextIdx] = 0
			}
			prev[nextIdx] = prev[nextIdx] | 2
			q = append(q, next)
		}
		next = NavPos{cur.coordinate, (cur.direction + 1) % 4}
		alt = dist[cur.getIndex(&matrix)] + 1000
		nextIdx = next.getIndex(&matrix)
		distNext = dist[nextIdx]
		if alt <= distNext {
			dist[nextIdx] = alt
			if alt < distNext {
				prev[nextIdx] = 0
			}
			prev[nextIdx] = prev[nextIdx] | 4
			q = append(q, next)
		}

	}

	minCost := math.MaxInt
	destinations := make([]NavPos, 0, 1)

	for i, _ := range directions {
		dest := NavPos{to, i}
		cost := dist[dest.getIndex(&matrix)]
		if cost <= minCost {
			if cost == minCost {
				destinations = make([]NavPos, 0, 1)
			}
			destinations = append(destinations, dest)
			minCost = cost
		}
	}

	counted := make(map[NavPos]bool)
	countTiles(&counted, &prev, destinations[0], &matrix, from)
	tileCount := 0
	counted2 := make(map[Vector]bool)

	for np, _ := range counted {
		if _, exist := counted2[np.coordinate]; exist {
			continue
		}
		counted2[np.coordinate] = true
		tileCount++
	}

	return minCost, tileCount
}

func countTiles(counted *map[NavPos]bool, prev *[]int, from NavPos, matrix *utils.Matrix[rune], start NavPos) {
	if _, exist := (*counted)[from]; exist {
		return
	}
	children := (*prev)[from.getIndex(matrix)]
	(*counted)[from] = true
	if from == start {
		return
	}

	if children&1 > 0 {
		next := NavPos{from.coordinate.subtract(directions[from.direction]), from.direction}
		countTiles(counted, prev, next, matrix, start)
	}
	if children&2 > 0 {
		next := NavPos{from.coordinate, (from.direction + 1) % 4}
		countTiles(counted, prev, next, matrix, start)
	}
	if children&4 > 0 {
		next := NavPos{from.coordinate, (from.direction + 3) % 4}
		countTiles(counted, prev, next, matrix, start)
	}
}

type Vector struct{ x, y int }

func (v1 *Vector) add(v2 Vector) Vector {
	return Vector{v1.x + v2.x, v1.y + v2.y}
}

func (v1 *Vector) subtract(v2 Vector) Vector {
	return Vector{v1.x - v2.x, v1.y - v2.y}
}

type NavPos struct {
	coordinate Vector
	direction  int
}

func (np NavPos) getIndex(matrix *utils.Matrix[rune]) int {
	return np.direction*(matrix.Width*matrix.Height) + (np.coordinate.x*matrix.Height + np.coordinate.y)
}
