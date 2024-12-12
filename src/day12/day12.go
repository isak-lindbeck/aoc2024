package day12

import (
	"github.com/isak-lindbeck/aoc2024/src/utils"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	matrix := utils.RuneMatrix(input)
	visited := utils.NewMatrixWithDefaultValue(matrix.Width, matrix.Height, false)
	for x, y := range matrix.Keys() {
		edges, area, corners := visitPlot(matrix, visited, x, y)
		ans1 += edges * area
		ans2 += corners * area
	}

	return ans1, ans2
}

func visitPlot(matrix utils.Matrix[rune], visited utils.Matrix[bool], x, y int) (int, int, int) {
	if visited.Get(x, y) {
		return 0, 0, 0
	}
	visited.Set(x, y, true)
	current := matrix.Get(x, y)
	edges := 0
	area := 1
	corners := 0

	// Orthogonal
	hasUpNeighbour := matrix.GetSafe(x, y-1, '.') == current    // ↑
	hasDownNeighbour := matrix.GetSafe(x, y+1, '.') == current  // ↓
	hasLeftNeighbour := matrix.GetSafe(x-1, y, '.') == current  // ←
	hasRightNeighbour := matrix.GetSafe(x+1, y, '.') == current // →
	// Diagonals
	hasUpRightNeighbour := matrix.GetSafe(x+1, y-1, '.') == current   // ↗
	hasDownRightNeighbour := matrix.GetSafe(x+1, y+1, '.') == current // ↘
	hasDownLeftNeighbour := matrix.GetSafe(x-1, y+1, '.') == current  // ↙
	hasUpLeftNeighbour := matrix.GetSafe(x-1, y-1, '.') == current    // ↖

	if !hasUpNeighbour {
		edges++
	} else {
		neighboursEdges, neighboursArea, neighboursCorners := visitPlot(matrix, visited, x, y-1)
		edges += neighboursEdges
		area += neighboursArea
		corners += neighboursCorners
	}
	if !hasDownNeighbour {
		edges++
	} else {
		neighboursEdges, neighboursArea, neighboursCorners := visitPlot(matrix, visited, x, y+1)
		edges += neighboursEdges
		area += neighboursArea
		corners += neighboursCorners
	}
	if !hasLeftNeighbour {
		edges++
	} else {
		neighboursEdges, neighboursArea, neighboursCorners := visitPlot(matrix, visited, x-1, y)
		edges += neighboursEdges
		area += neighboursArea
		corners += neighboursCorners
	}
	if !hasRightNeighbour {
		edges++
	} else {
		neighboursEdges, neighboursArea, neighboursCorners := visitPlot(matrix, visited, x+1, y)
		edges += neighboursEdges
		area += neighboursArea
		corners += neighboursCorners
	}

	cornerCases := []bool{
		// Outer Corners
		!hasLeftNeighbour && !hasUpNeighbour,
		!hasUpNeighbour && !hasRightNeighbour,
		!hasRightNeighbour && !hasDownNeighbour,
		!hasDownNeighbour && !hasLeftNeighbour,
		// Inner Corners
		hasUpNeighbour && hasRightNeighbour && !hasUpRightNeighbour,
		hasDownNeighbour && hasRightNeighbour && !hasDownRightNeighbour,
		hasDownNeighbour && hasLeftNeighbour && !hasDownLeftNeighbour,
		hasUpNeighbour && hasLeftNeighbour && !hasUpLeftNeighbour,
	}
	for _, cornerCase := range cornerCases {
		if cornerCase {
			corners++
		}
	}

	return edges, area, corners
}
