package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	minSteps := make([]int, len(nums))
	for i := 1; i < len(minSteps); i++ {
		minSteps[i] = len(nums) + 100
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
			if minSteps[i]+1 < minSteps[j] {
				minSteps[j] = minSteps[i] + 1
			}
		}
	}
	fmt.Println(minSteps)
	return minSteps[len(nums)-1]
}
