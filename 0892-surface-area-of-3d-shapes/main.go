package main

import "fmt"

func main() {
	fmt.Println(surfaceArea([][]int{[]int{2}}))
	fmt.Println(surfaceArea([][]int{[]int{1, 2}, []int{3, 4}}))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func surfaceArea(grid [][]int) int {
	res := 0
	// 横着扫一遍
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 第一行
			if i == 0 {
				res += grid[i][j]
				continue
			}

			// 中间
			res += abs(grid[i][j] - grid[i-1][j])
		}
		// 最后一行
		if i == len(grid)-1 {
			for j := 0; j < len(grid[0]); j++ {
				res += grid[i][j]
			}
		}
	}

	// 竖着扫一遍
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			// grid[j][i]
			// 第一列
			if i == 0 {
				res += grid[j][i]
				continue
			}
			// 中间
			res += abs(grid[j][i] - grid[j][i-1])
		}
		// 最后一列
		if i == len(grid[0])-1 {
			for j := 0; j < len(grid); j++ {
				res += grid[j][i]
			}
		}
	}

	// 天花板 && 地下室
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				res += 2
			}
		}
	}
	return res
}
