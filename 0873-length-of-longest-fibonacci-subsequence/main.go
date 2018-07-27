package main

func lenLongestFibSubseq(A []int) int {
	if len(A) < 3 {
		return 0
	}

	var indexMap = map[int]int{}
	for i := 0; i < len(A); i++ {
		indexMap[A[i]] = i
	}

	var maxLen = 2

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			currentLen := 2
			i, j := i, j
			for {
				sum := A[i] + A[j]
				if indexMap[sum] > j {
					currentLen++
					if currentLen > maxLen {
						maxLen = currentLen
					}
					i = j
					j = indexMap[sum]
				} else {
					break
				}
			}
		}
	}

	if maxLen == 2 {
		return 0
	}
	return maxLen
}
