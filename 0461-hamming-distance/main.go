package main

func hammingDistance(x int, y int) int {
	// 将 x 与 y 不同的位置 1
	// 所以就是异或的逻辑
	xorResult := x ^ y
	cnt := 0
	for xorResult > 0 {
		if xorResult%2 == 1 {
			cnt++
		}
		xorResult /= 2
	}
	return cnt
}
