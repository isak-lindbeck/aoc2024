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

func ParseSlice(s []string) []int {
	ints := make([]int, len(s))
	for i, v := range s {
		ints[i] = Parse(v)
	}
	return ints
}

func Parse(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(s + " is not a number")
	}
	return i
}

func FromBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Pow(n, exponent int) int {
	result := 1
	for i := 1; i <= exponent; i++ {
		result *= n
	}
	return result
}
