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
		region := visitPlot(matrix, visited, x, y)
		ans1 += region.area * region.edges
		ans2 += region.area * region.corners
	}

	return ans1, ans2
}

func visitPlot(matrix utils.Matrix[rune], visited utils.Matrix[bool], x, y int) Region {
	if visited.Get(x, y) {
		return Region{0, 0, 0}
	}
	visited.Set(x, y, true)
	region := Region{0, 1, 0}

	current := matrix.Get(x, y)
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

	if hasUpNeighbour {
		region.merge(visitPlot(matrix, visited, x, y-1))
	}
	if hasDownNeighbour {
		region.merge(visitPlot(matrix, visited, x, y+1))
	}
	if hasLeftNeighbour {
		region.merge(visitPlot(matrix, visited, x-1, y))
	}
	if hasRightNeighbour {
		region.merge(visitPlot(matrix, visited, x+1, y))
	}

	for _, hasNeighbour := range []bool{hasUpNeighbour, hasDownNeighbour, hasLeftNeighbour, hasRightNeighbour} {
		if !hasNeighbour {
			region.edges++
		}
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
			region.corners++
		}
	}

	return region
}

type Region struct {
	edges, area, corners int
}

func (r *Region) merge(other Region) {
	r.edges += other.edges
	r.area += other.area
	r.corners += other.corners
}
