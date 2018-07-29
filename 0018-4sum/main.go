package main

import "sort"

func exist(nums []int, numsCollection [][]int) bool {
	sort.Ints(nums)

	for _, n := range numsCollection {
		sort.Ints(n)
		equal := true
		for i := 0; i < len(n); i++ {
			if nums[i] != n[i] {
				equal = false
				break
			}
		}

		if equal {
			return true
		}
	}
	return false
}

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k]+nums[l] == target {
						current := []int{nums[i], nums[j], nums[k], nums[l]}
						if !exist(current, result) {
							result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
						}
					}
				}
			}
		}
	}
	return result
}
