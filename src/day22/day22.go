package day22

import (
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"strings"
)

const side4dSlice = 19 // Covers values -9 -> +9
var sequenceSums = create4dSlice[int]()

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	split := strings.Split(input, "\n")

	for _, s := range split {
		secretNumber := ints.Parse(s)
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
					sequenceSums[idx] += nexVal
					isSet[idx] = true
				}
			}
			sequenceOffset++
			secretNumber = next
		}

		ans1 += secretNumber
	}
	ans2 = getMaxValue(sequenceSums)

	return ans1, ans2
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

func getMaxValue(count []int) int {
	maxVal := 0
	for i := range count {
		if count[i] > maxVal {
			maxVal = count[i]
		}
	}
	return maxVal
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
