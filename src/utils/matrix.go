package utils

import (
	"strings"
)

type Coord struct {
	X, Y int
}

type Matrix[T comparable] struct {
	Width  int
	Height int
	data   []T
}

func NewMatrix[T comparable](width, height int) Matrix[T] {
	return Matrix[T]{
		Width:  width,
		Height: height,
		data:   make([]T, width*height),
	}
}

func NewMatrixWithDefaultValue[T comparable](width, height int, defaultValue T) Matrix[T] {
	data := make([]T, width*height)
	for i := range data {
		data[i] = defaultValue
	}
	return Matrix[T]{
		Width:  width,
		Height: height,
		data:   data,
	}
}

func CloneMatrix[T comparable](m Matrix[T]) Matrix[T] {
	data := make([]T, m.Width*m.Height)
	copy(data, m.data)
	return Matrix[T]{
		Width:  m.Width,
		Height: m.Height,
		data:   data,
	}
}

func RuneMatrix(input string) Matrix[rune] {
	input = strings.TrimSuffix(input, "\n")
	width := strings.IndexRune(input, '\n')
	height := strings.Count(input, "\n") + 1
	input = strings.ReplaceAll(input, "\n", "")

	matrix := NewMatrix[rune](height, width)
	for i, r := range input {
		x := i % width
		y := i / height
		matrix.Set(x, y, r)
	}
	return matrix
}

func (m *Matrix[T]) Get(x, y int) T {
	return m.data[x*m.Height+y]
}

func (m *Matrix[T]) GetSafe(x, y int, defaultValue T) T {
	if x < 0 || y < 0 || x >= m.Width || y >= m.Height {
		return defaultValue
	}
	return m.Get(x, y)
}

func (m *Matrix[T]) Set(x, y int, value T) {
	m.data[x*m.Height+y] = value
}

func (m *Matrix[T]) Keys() func(yield func(int, int) bool) {
	return func(yield func(int, int) bool) {
		for y := 0; y < m.Height; y++ {
			for x := 0; x < m.Width; x++ {
				if !yield(x, y) {
					return
				}
			}
		}
	}
}

func (m *Matrix[T]) GetCoordinates(v T) (int, int) {
	for x, y := range m.Keys() {
		if v == m.Get(x, y) {
			return x, y
		}
	}
	return -1, -1
}
