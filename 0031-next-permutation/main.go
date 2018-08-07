package main

import "sort"

func main() {
}

func nextPermutation(nums []int) {
	found := false
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			found = true
			swapIdx := i
			for j := i + 1; j < len(nums); j++ {
				if nums[j] > nums[i-1] && nums[j] < nums[i] {
					swapIdx = j
				}
			}
			nums[i-1], nums[swapIdx] = nums[swapIdx], nums[i-1]
			sort.Ints(nums[i:])
			break
		}
	}

	if found == false {
		for i := 0; i < len(nums)/2; i++ {
			nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
		}
	}
}
