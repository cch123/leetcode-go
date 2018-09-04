package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}))
}

type point struct {
	x int
	y int
}

func dfs(res [][]int, grid [][]int, next []point) {
	if len(next) == 0 {
		return
	}

	// 计算现在的
	if len(next) >= 1 {
		for i := 0; i < len(next); i++ {
			// left
			x, y := next[i].x, next[i].y-1
			if y >= 0 && res[x][y]+grid[x][y+1] < res[x][y+1] {
				res[x][y+1] = res[x][y] + grid[x][y+1]
			}
			// up
			x, y = next[i].x-1, next[i].y
			if x >= 0 && res[x][y]+grid[x+1][y] < res[x+1][y] {
				res[x+1][y] = res[x][y] + grid[x+1][y]
			}
		}
	}

	newNext := []point{}
	distinctMap := map[point]bool{}
	for i := 0; i < len(next); i++ {
		// right
		x, y := next[i].x, next[i].y+1
		_, ok := distinctMap[point{x, y}]
		if y < len(grid[0]) && !ok {
			newNext = append(newNext, point{x, y})
			distinctMap[point{x, y}] = true
		}
		// down
		x, y = next[i].x+1, next[i].y
		_, ok = distinctMap[point{x, y}]
		if x < len(grid) && !ok {
			newNext = append(newNext, point{x, y})
			distinctMap[point{x, y}] = true
		}
	}
	dfs(res, grid, newNext)
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	res := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		res[i] = make([]int, len(grid[0]))
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = 99999999999999999
		}
	}
	res[0][0] = grid[0][0]
	dfs(res, grid, []point{point{0, 0}})
	return res[len(grid)-1][len(grid[0])-1]
}
