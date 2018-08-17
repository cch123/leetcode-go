package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getResult(stack []int) []string {
	result := []string{}
	n := len(stack)
	for i := 0; i < len(stack); i++ {
		line := make([]byte, n)
		for j := 0; j < len(line); j++ {
			line[j] = '.'
		}
		line[stack[i]] = 'Q'
		result = append(result, string(line))
	}

	return result
}

func traverse(parentStack []int, currentLevel int, maxLevel int, res *[][]string) {
	//fmt.Println(parentStack, currentLevel, maxLevel)
	if currentLevel == maxLevel {
		currentMatrix := getResult(parentStack)
		*res = append(*res, currentMatrix)
		return
	}

	for i := 0; i < maxLevel; i++ {
		valid := true
		for j := 0; j < len(parentStack); j++ {
			r1, c1 := j, parentStack[j]
			r2, c2 := currentLevel, i
			if c1 == c2 || abs(r1-r2) == abs(c1-c2) {
				// valid
				valid = false
			}
		}
		if valid {
			currentStack := append([]int{}, parentStack...)
			currentStack = append(currentStack, i)
			traverse(currentStack, currentLevel+1, maxLevel, res)
		}
	}
}

func solveNQueens(n int) [][]string {
	res := [][]string{}
	traverse([]int{}, 0, n, &res)
	return res
}
