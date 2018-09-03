package main

import "fmt"

func main() {
	fmt.Println(subarrayBitwiseORs([]int{0}))
	fmt.Println(subarrayBitwiseORs([]int{1, 1, 2}))
	fmt.Println(subarrayBitwiseORs([]int{1, 2, 4}))
}

func subarrayBitwiseORs(A []int) int {
	bitwiseArr := make([]int, len(A))
	resMap := map[int]bool{}

	// 起始位置
	for i := 0; i < len(A); i++ {
		// 起始位置
		for j := i; j < len(A); j++ {
			if j-i == 0 {
				bitwiseArr[j] = A[j]
			} else {
				bitwiseArr[j] = bitwiseArr[j-1] | A[j]
			}
			if resMap[bitwiseArr[j]] == false {
				resMap[bitwiseArr[j]] = true
			}
		}
	}
	return len(resMap)
}
