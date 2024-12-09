package day09

import (
	"github.com/isak-lindbeck/aoc2024/src/ints"
	"slices"
	"strings"
)

func Run(input string) (int, int) {
	ans1 := 0
	ans2 := 0

	input = strings.TrimSuffix(input, "\n")
	diskChunks := make([]int, len(input))
	for i, r := range input {
		diskChunks[i] = ints.Parse(string(r))
	}

	ans1 = runPartOne(slices.Clone(diskChunks))

	ans2 = runPartTwo(diskChunks)

	return ans1, ans2
}

func runPartOne(diskChunks []int) int {
	sum := 0
	lastIndex := len(diskChunks) - 1
	k := 0
	for index := 0; index <= lastIndex; index++ {
		for j := 0; j < diskChunks[index]; j++ {
			if index%2 == 0 {
				sum += k * (index / 2)
			} else {
				if diskChunks[lastIndex] == 0 {
					lastIndex -= 2
				}
				diskChunks[lastIndex]--
				sum += k * (lastIndex / 2)
			}
			k++
		}
	}

	return sum
}

func runPartTwo(diskChunks []int) int {
	sum := 0
	usedFreeSpaces := slices.Repeat([]int{0}, len(diskChunks))
	totalLengthBefore := make([]int, len(diskChunks))
	total := 0
	for i := 1; i < len(diskChunks); i++ {
		total += diskChunks[i-1]
		totalLengthBefore[i] = total
	}

	lastIndex := len(diskChunks) - 1
	for lastIndex > 0 {
		movedId := lastIndex / 2
		movingChunk := diskChunks[lastIndex]
		moved := false
		for i := 1; i < lastIndex; i += 2 {
			freeChunk := diskChunks[i]
			if freeChunk >= movingChunk {
				offset := usedFreeSpaces[i] + totalLengthBefore[i]
				for l := range movingChunk {
					sum += (l + offset) * movedId
				}
				usedFreeSpaces[i] += movingChunk
				diskChunks[i] -= movingChunk
				moved = true
				break
			}
		}
		if !moved {
			offset := totalLengthBefore[lastIndex]
			for l := range movingChunk {
				sum += (l + offset) * movedId
			}
		}
		lastIndex -= 2
	}
	return sum
}
