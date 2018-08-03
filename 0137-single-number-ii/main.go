package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for j := 0; j < 64; j++ {
		sum := 0
		for i := 0; i < len(nums); i++ {
			if (nums[i] >> uint(j) & 1) > 0 {
				sum++
			}
		}
		if sum%3 > 0 {
			res |= 1 << uint(j)
		}
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}
