package day22

import (
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"strings"
)

const side4dSlice = 19 // Covers values -9 -> +9

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	split := strings.Split(input, "\n")

	sequenceSum := create4dSlice(0)
	for _, s := range split {
		secretNumber := ints.Parse(s)
		sequence := []int{0, 0, 0, 0}
		idx := 0
		isSet := create4dSlice(false)
		for unused := 0; unused < 2000; unused++ {
			next := processNumber(secretNumber)
			nexVal := next % 10
			change := nexVal - (secretNumber % 10)
			sequence[idx%4] = change
			if idx > 3 {
				i := sequence[(idx+0)%4] + 9
				j := sequence[(idx+1)%4] + 9
				k := sequence[(idx+2)%4] + 9
				l := sequence[(idx+3)%4] + 9
				if !isSet[i][j][k][l] {
					sequenceSum[i][j][k][l] += nexVal
					isSet[i][j][k][l] = true
				}
			}
			idx++
			secretNumber = next
		}

		ans1 += secretNumber
	}
	ans2 = getMaxValue(sequenceSum)

	return ans1, ans2
}

func processNumber(secretNumber int) int {
	secretNumber = ((secretNumber * 64) ^ secretNumber) % 16777216
	secretNumber = ((secretNumber / 32) ^ secretNumber) % 16777216
	secretNumber = ((secretNumber * 2048) ^ secretNumber) % 16777216
	return secretNumber
}

func create4dSlice[T any](defaultValue T) [][][][]T {
	count := make([][][][]T, side4dSlice)
	for i := range side4dSlice {
		count[i] = make([][][]T, side4dSlice)
		for j := range side4dSlice {
			count[i][j] = make([][]T, side4dSlice)
			for k := range side4dSlice {
				count[i][j][k] = make([]T, side4dSlice)
				for l := range side4dSlice {
					count[i][j][k][l] = defaultValue
				}
			}
		}
	}
	return count
}

func getMaxValue(count [][][][]int) int {
	maxVal := 0
	for i := range side4dSlice {
		for j := range side4dSlice {
			for k := range side4dSlice {
				for l := range side4dSlice {
					if count[i][j][k][l] > maxVal {
						maxVal = count[i][j][k][l]
					}
				}
			}
		}
	}
	return maxVal
}
