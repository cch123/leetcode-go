package main

import "fmt"

func main() {
	fmt.Println(countBits(5))
}

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	if num == 1 {
		return []int{0, 1}
	}
	res := make([]int, num+1)
	res[0], res[1] = 0, 1
	round := 1
outer:
	for {
		round = round * 2
		for i := round; i < round*2; i++ {
			if i > num {
				break outer
			}
			res[i] = res[i-round] + 1
		}
	}
	return res
}
