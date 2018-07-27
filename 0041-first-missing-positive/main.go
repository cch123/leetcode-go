package main

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	var arr = make([]bool, len(nums)+1)
	var idx = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 && nums[i] < len(arr) {
			arr[nums[i]] = true
		}
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] == false {
			idx = i
			break
		} else {
			idx = i + 1
		}
	}

	return idx
}
