package main

import "fmt"

func projectionArea(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	elemCnt := len(grid) * len(grid[0])
	colMaxSum := 0
	rowMaxSum := 0
	for i := 0; i < len(grid); i++ {
		rowIMax := 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > rowIMax {
				rowIMax = grid[i][j]
			}
			if grid[i][j] == 0 {
				elemCnt--
			}
		}
		rowMaxSum += rowIMax
	}
	for i := 0; i < len(grid[0]); i++ {
		colIMax := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > colIMax {
				colIMax = grid[j][i]
			}
		}
		colMaxSum += colIMax
	}
	return elemCnt + colMaxSum + rowMaxSum

}

func main() {
	fmt.Println(projectionArea([][]int{[]int{1, 0}, []int{0, 2}}))
}
