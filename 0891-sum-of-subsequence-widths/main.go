package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sumSubseqWidths([]int{2, 1, 3}))
}

func pow(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}

	if n%2 == 0 {
		mid := pow(n/2) % 1000000007
		return (mid * mid) % 1000000007
	}
	mid := pow(n/2) % 1000000007
	return (2 * mid * mid) % 1000000007
}

func sumSubseqWidths(A []int) int {
	res := 0
	sort.Ints(A)
	for i := 0; i < len(A); i++ {
		res = res + (pow(i)-pow(len(A)-i-1))*A[i]
		res = res % 1000000007
		//println(i, A[i], res, pow(i), pow(len(A)-i-1))
	}
	return res
}
