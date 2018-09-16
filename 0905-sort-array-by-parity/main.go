package main

import "fmt"

func main() {
	arr := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(arr))
}

func sortArrayByParity(A []int) []int {
	pos := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			A[pos], A[i] = A[i], A[pos]
			pos++
		}
	}
	return A
}
