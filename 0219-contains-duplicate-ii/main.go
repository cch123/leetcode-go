package main

func containsNearbyDuplicate(nums []int, k int) bool {
	var numMap = map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; ok {
			if i-numMap[nums[i]] <= k {
				return true
			}
		}
		numMap[nums[i]] = i
	}
	return false
}
