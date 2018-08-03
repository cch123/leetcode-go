package main

import "fmt"

func rowResultWillBeZero(matrix [][]int, row int) bool {
	r := matrix[row]
	for i := 0; i < len(r); i++ {
		if r[i] == 0 {
			return true
		}
	}
	return false
}

func colResultWillBeZero(matrix [][]int, col int) bool {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][col] == 0 {
			return true
		}
	}
	return false
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	markRow, markCol := 0, -1
	colCnt := len(matrix[0])
	firstRow := matrix[0]

	for i := 0; i < colCnt; i++ {
		if firstRow[i] != 0 {
			markCol = i
			break
		}
	}

	if markCol == -1 {
		// mark all to zero
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
		return
	}
	markRowIsZero := rowResultWillBeZero(matrix, markRow)
	markColIsZero := colResultWillBeZero(matrix, markCol)

	// traverse

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == markCol {
				continue
			}
			if matrix[i][j] == 0 {
				// set mark pos = 0
				matrix[markRow][j] = 0
				matrix[i][markCol] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == markCol {
				continue
			}
			if matrix[markRow][j] == 0 || matrix[i][markCol] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// mark markrow and markcol
	if markRowIsZero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	if markColIsZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][markCol] = 0
		}
	}

}

func main() {
	var m = [][]int{[]int{0}}
	setZeroes(m)
	fmt.Println(m)
	m = [][]int{[]int{0, 1, 2, 0}, []int{3, 4, 5, 2}, []int{1, 3, 1, 5}}
	setZeroes(m)
	fmt.Println(m)
	m = [][]int{[]int{0}, []int{1}}
	setZeroes(m)
	fmt.Println(m)
}
