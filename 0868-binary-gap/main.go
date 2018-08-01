package main

import (
	"strconv"
)

func binaryGap(N int) int {
	nBinaryStr := strconv.FormatInt(int64(N), 2)
	maxGap := 0
	currentGap := 0
	for i := 1; i < len(nBinaryStr); i++ {
		currentGap++
		if nBinaryStr[i] == '0' {
			continue
		}

		if nBinaryStr[i] == '1' {
			if currentGap > maxGap {
				maxGap = currentGap
			}
			currentGap = 0
		}
	}

	return maxGap
}

func main() {
	println(binaryGap(22))
	println(binaryGap(5))
	println(binaryGap(6))
}
