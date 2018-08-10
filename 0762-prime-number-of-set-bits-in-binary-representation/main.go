package main

import "math"

func main() {
}

func countPrimeSetBits(L int, R int) int {
	cnt := 0
	for i := L; i <= R; i++ {
		bitCnt := countBits(i)
		if isPrime(bitCnt) {
			cnt++
		}
	}
	return cnt
}

func countBits(n int) int {
	cnt := 0
	for n > 0 {
		if n&1 > 0 {
			cnt++
		}
		n = n >> 1
	}
	return cnt
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
