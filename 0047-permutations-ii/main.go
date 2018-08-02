package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	copyNum := append([]int{}, nums...)

	res := [][]int{}
	res = append(res, copyNum)
	for {
		canBeBigger := false
		for i := len(nums) - 1; i >= 1; i-- {
			if nums[i] > nums[i-1] {
				canBeBigger = true
				// 从 i 向后找最小的 > nums[i-1] 的数
				// 排序 nums[i:]
				for j := i; j < len(nums); j++ {
					if nums[j] < nums[i] && nums[j] > nums[i-1] {
						nums[j], nums[i] = nums[i], nums[j]
					}
				}
				nums[i], nums[i-1] = nums[i-1], nums[i]
				sort.Ints(nums[i:])
				internalRes := append([]int{}, nums...)
				res = append(res, internalRes)
				break
			}
		}
		if !canBeBigger {
			break
		}
	}
	return res
}

func main() {}
