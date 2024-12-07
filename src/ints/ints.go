package ints

import (
	"log"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Parse(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(s + " is not a number")
	}
	return i
}

func Pow(n, m int) int {
	result := 1
	for i := 1; i <= m; i++ {
		result *= n
	}
	return result
}
