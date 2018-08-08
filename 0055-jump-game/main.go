package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{0}))
}

func canJump(nums []int) bool {

	var can = make([]bool, len(nums))
	can[0] = true

	for i := 0; i < len(nums); i++ {
		if can[i] != true {
			continue
		}

		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
			can[j] = true
		}
	}

	fmt.Println(can)

	return can[len(nums)-1]
}
