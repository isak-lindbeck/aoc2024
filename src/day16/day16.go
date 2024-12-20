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
	ans1, ans2 = Dijkstra(matrix, Node{Vector{fromX, fromY}, 1}, Vector{toX, toY})

	return ans1, ans2
}

func Dijkstra(matrix utils.Matrix[rune], from Node, to Vector) (int, int) {
	size := matrix.Width * matrix.Height * 4
	dist := slices.Repeat([]int{math.MaxInt - 1}, size)
	prev := make([]Node, size)

	dist[from.getIndex(&matrix)] = 0

	minToCost := math.MaxInt
	queue := utils.NewQueue(make([]Node, size))
	queue.PushFront(from)
	for true {
		cur, exists := queue.Pop()
		if !exists {
			break
		}
		curDist := dist[cur.getIndex(&matrix)]
		if cur.coordinate == to && curDist < minToCost {
			minToCost = curDist
		}

		next := Node{cur.coordinate.add(directions[cur.direction]), cur.direction}
		canMove := matrix.Get(next.coordinate.x, next.coordinate.y) == '.'
		if canMove {
			alt := curDist + 1
			nextIdx := next.getIndex(&matrix)
			if alt <= dist[nextIdx] && alt < minToCost {
				dist[nextIdx] = alt
				prev[nextIdx] = cur
				queue.PushFront(next)
			}
		}
		next = Node{cur.coordinate, (cur.direction + 3) % 4}
		alt := curDist + 1000
		nextIdx := next.getIndex(&matrix)
		if alt <= dist[nextIdx] && alt < minToCost {
			dist[nextIdx] = alt
			prev[nextIdx] = cur
			queue.PushBack(next)
		}
		next = Node{cur.coordinate, (cur.direction + 1) % 4}
		alt = curDist + 1000
		nextIdx = next.getIndex(&matrix)
		if alt <= dist[nextIdx] && alt < minToCost {
			dist[nextIdx] = alt
			prev[nextIdx] = cur
			queue.PushBack(next)
		}

	}

	minCost := math.MaxInt
	destinations := make([]Node, 0, 1)

	for i, _ := range directions {
		dest := Node{to, i}
		cost := dist[dest.getIndex(&matrix)]
		if cost <= minCost {
			if cost == minCost {
				destinations = make([]Node, 0, 1)
			}
			destinations = append(destinations, dest)
			minCost = cost
		}
	}

	counted := make(map[Node]bool)
	countTiles(&counted, &prev, &dist, destinations[0], &matrix, from)
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

func countTiles(counted *map[Node]bool, prev *[]Node, dist *[]int, from Node, matrix *utils.Matrix[rune], start Node) {
	if _, exist := (*counted)[from]; exist {
		return
	}
	(*counted)[from] = true
	if from == start {
		return
	}

	neighbors := []Node{
		{from.coordinate.subtract(directions[from.direction]), from.direction},
		{from.coordinate, (from.direction + 3) % 4},
		{from.coordinate, (from.direction + 1) % 4},
	}

	pp := (*prev)[from.getIndex(matrix)]
	previous := pp.getIndex(matrix)
	lowestCost := (*dist)[previous]
	for _, neighbor := range neighbors {
		nCost := (*dist)[neighbor.getIndex(matrix)]
		if nCost == lowestCost || nCost == lowestCost-999 {
			countTiles(counted, prev, dist, neighbor, matrix, start)
		}
	}
}

type Vector struct{ x, y int }

func (v1 *Vector) add(v2 Vector) Vector {
	return Vector{v1.x + v2.x, v1.y + v2.y}
}

func (v1 *Vector) subtract(v2 Vector) Vector {
	return Vector{v1.x - v2.x, v1.y - v2.y}
}

type Node struct {
	coordinate Vector
	direction  int
}

func (np Node) getIndex(matrix *utils.Matrix[rune]) int {
	return np.direction*(matrix.Width*matrix.Height) + (np.coordinate.x*matrix.Height + np.coordinate.y)
}
