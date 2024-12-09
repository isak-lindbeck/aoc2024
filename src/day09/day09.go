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
	ans := 0
	lastChunkIdx := len(diskChunks) - 1
	blockIdx := 0
	for chunkIdx := 0; chunkIdx <= lastChunkIdx; chunkIdx++ {
		for range diskChunks[chunkIdx] {
			isFreeChunk := chunkIdx%2 == 0
			if isFreeChunk {
				fileId := chunkIdx / 2
				ans += blockIdx * fileId
			} else {
				if diskChunks[lastChunkIdx] == 0 {
					lastChunkIdx -= 2
				}
				diskChunks[lastChunkIdx]--
				fileId := lastChunkIdx / 2
				ans += blockIdx * fileId
			}
			blockIdx++
		}
	}

	return ans
}

func runPartTwo(diskChunks []int) int {
	ans := 0
	usedSpaceIn := make([]int, len(diskChunks))
	totalLengthBefore := make([]int, len(diskChunks))
	total := 0
	for i := 1; i < len(diskChunks); i++ {
		total += diskChunks[i-1]
		totalLengthBefore[i] = total
	}

	firstFreeIdxOfSize := slices.Repeat([]int{1}, 10)
	for movingFileIdx := len(diskChunks) - 1; movingFileIdx > 0; movingFileIdx -= 2 {
		fileCanMove := false
		movingFileSize := diskChunks[movingFileIdx]
		freeChunkIdx := firstFreeIdxOfSize[movingFileSize]
		for freeChunkIdx < movingFileIdx {
			if diskChunks[freeChunkIdx] >= movingFileSize {
				fileCanMove = true
				break
			}
			freeChunkIdx += 2
		}
		fileId := movingFileIdx / 2
		if fileCanMove {
			offset := usedSpaceIn[freeChunkIdx] + totalLengthBefore[freeChunkIdx]
			for blockIdx := range movingFileSize {
				ans += (blockIdx + offset) * fileId
			}
			usedSpaceIn[freeChunkIdx] += movingFileSize
			diskChunks[freeChunkIdx] -= movingFileSize

			firstFreeIdxOfSize[movingFileSize] = freeChunkIdx
		} else {
			offset := totalLengthBefore[movingFileIdx]
			for blockIdx := range movingFileSize {
				ans += (blockIdx + offset) * fileId
			}
		}
	}
	return ans
}
