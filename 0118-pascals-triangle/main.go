package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	var res = make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		internalArr := make([]int, i+1)
		internalArr[0] = 1
		internalArr[len(internalArr)-1] = 1
		for j := 1; j < len(internalArr)-1; j++ {
			internalArr[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = internalArr
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}
