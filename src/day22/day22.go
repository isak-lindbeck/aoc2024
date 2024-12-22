package day22

import (
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"slices"
	"strings"
	"sync/atomic"
)

const side4dSlice = 19 // Covers values -9 -> +9
var sequenceSums = create4dSlice[uint32]()

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	lines := strings.Fields(input)

	c := make(chan int)
	for _, s := range lines {
		secretNumber := ints.Parse(s)
		go calculateSecretNums(secretNumber, c)
	}
	for range lines {
		ans1 += <-c
	}
	ans2 = int(slices.Max(sequenceSums))

	return ans1, ans2
}

func calculateSecretNums(secretNumber int, c chan int) {
	sequence := []int{0, 0, 0, 0}
	sequenceOffset := 0
	isSet := create4dSlice[bool]()
	for unused := 0; unused < 2000; unused++ {
		next := processNumber(secretNumber)
		nexVal := next % 10
		change := nexVal - (secretNumber % 10)
		sequence[sequenceOffset%4] = change
		if sequenceOffset > 3 {
			idx := toIdx(sequence, sequenceOffset)
			if !isSet[idx] {
				atomic.AddUint32(&sequenceSums[idx], uint32(nexVal))
				isSet[idx] = true
			}
		}
		sequenceOffset++
		secretNumber = next
	}
	c <- secretNumber
}

func processNumber(secretNumber int) int {
	secretNumber = ((secretNumber * 64) ^ secretNumber) % 16777216
	secretNumber = ((secretNumber / 32) ^ secretNumber) % 16777216
	secretNumber = ((secretNumber * 2048) ^ secretNumber) % 16777216
	return secretNumber
}

func create4dSlice[T any]() []T {
	return make([]T, side4dSlice*side4dSlice*side4dSlice*side4dSlice)
}

func toIdx(sequence []int, idx int) int {
	i := sequence[(idx+0)%4] + 9 // Offset -9 -> 0
	j := sequence[(idx+1)%4] + 9
	k := sequence[(idx+2)%4] + 9
	l := sequence[(idx+3)%4] + 9
	i = i * side4dSlice * side4dSlice * side4dSlice
	j = j * side4dSlice * side4dSlice
	k = k * side4dSlice
	return i + j + k + l
}
