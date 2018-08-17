package main

import "fmt"

func main() {
	fmt.Println(4, totalNQueens(4))
	fmt.Println(8, totalNQueens(8))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func traverse(parentStack []int, level int, n int, res *int) {
	//fmt.Println("level", level, "n", n, parentStack)
	if level == n {
		//fmt.Println(parentStack)
		*res++
	}

	for i := 0; i < n; i++ {
		valid := true
		for j := 0; j < len(parentStack); j++ {
			if parentStack[j] == i || abs(parentStack[j]-i) == abs(j-level) {
				valid = false
				break
			}
		}

		if valid {
			currentStack := append([]int{}, parentStack...)
			currentStack = append(currentStack, i)
			traverse(currentStack, level+1, n, res)
		}
	}
}

func totalNQueens(n int) int {
	var res int
	traverse([]int{}, 0, n, &res)
	return res
}
