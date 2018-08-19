package main

import "fmt"

func main() {
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2}))
}

func fairCandySwap(A []int, B []int) []int {
	var diff int
	var aMap, bMap = map[int]bool{}, map[int]bool{}
	var sumA, sumB int
	for i := 0; i < len(A); i++ {
		aMap[A[i]] = true
		sumA += A[i]
	}
	for i := 0; i < len(B); i++ {
		bMap[B[i]] = true
		sumB += B[i]
	}
	diff = sumA - sumB

	if diff%2 != 0 {
		return []int{}
	}

	for i := 0; i < len(A); i++ {
		if bMap[A[i]-diff/2] == true {
			return []int{A[i], A[i] - diff/2}
		}
	}
	return []int{}
}
