package main

func transpose(A [][]int) [][]int {
	if len(A) == 0 {
		return A
	}
	x := len(A)
	y := len(A[0])
	resultMatrix := make([][]int, y)
	for i := 0; i < len(resultMatrix); i++ {
		resultMatrix[i] = make([]int, x)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			resultMatrix[j][i] = A[i][j]
		}
	}
	return resultMatrix
}
