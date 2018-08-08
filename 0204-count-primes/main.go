package main

import (
	"math"
)

func main() {
	println(countPrimes(10))
	println(countPrimes(20))
	println(countPrimes(100))
}

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	var numArr = make([]int, n+1)
	numArr[0] = -1
	numArr[1] = -1
	for i := 2; i <= int(math.Sqrt(float64(n)))+1; i++ {
		for j := 2 * i; j <= n; j += i {
			numArr[j] = -1
		}
	}
	cnt := 0
	for i := 2; i < n; i++ {
		if numArr[i] != -1 {
			cnt++
		}
	}
	//fmt.Printf("%#v\n", numArr)
	return cnt
}
