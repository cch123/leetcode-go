package main

import (
	"fmt"
	"sort"
)

/*
@1 : 从右向左搜索，找到第一个 i，满足 nums[i] > nums[i-1]
@2 : 从 i 向右(包括i)搜索，找出满足 > nums[i-1] 条件的最小的 nums[j]
@3 : swap(nums[i-1], nums[j])
@4 : sort(nums[i:])
*/
func nextPermutation(nums []int) []int {
	// @1
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			// @2
			minMax, minMaxIdx := nums[i], i
			for j := i; j < len(nums); j++ {
				if nums[j] < minMax && nums[j] > nums[i-1] {
					minMax, minMaxIdx = nums[j], j
				}
			}

			// @3
			nums[i-1], nums[minMaxIdx] = nums[minMaxIdx], nums[i-1]
			// @4
			subSlice := nums[i:]
			sort.Slice(subSlice, func(x, y int) bool {
				return subSlice[x] < subSlice[y]
			})
			break
		}
	}
	return nums
}

func main() {
	fmt.Println(nextPermutation([]int{1, 2, 3, 4, 5}))
	fmt.Println(nextPermutation([]int{1, 2, 5, 4, 3}))
}
