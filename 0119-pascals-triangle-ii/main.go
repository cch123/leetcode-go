package main

import "fmt"

func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	}

	var res = make([][]int, rowIndex+1)
	res[0] = []int{1}
	for i := 1; i <= rowIndex; i++ {
		internalArr := make([]int, i+1)
		internalArr[0] = 1
		internalArr[len(internalArr)-1] = 1
		for j := 1; j < len(internalArr)-1; j++ {
			internalArr[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = internalArr
	}
	return res[len(res)-1]
}

func main() {
	fmt.Println(getRow(3))
	fmt.Println(getRow(5))
}
