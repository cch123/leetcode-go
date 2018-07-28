package main

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		mid := (len(A[i]) + 1) / 2
		for j := 0; j < mid; j++ {
			k := len(A[i]) - 1 - j
			A[i][j], A[i][k] = 1-A[i][k], 1-A[i][j]
		}

	}

	return A
}
