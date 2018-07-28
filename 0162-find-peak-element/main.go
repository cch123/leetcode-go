package main

var res = []int{}

func find(nums []int, startIdx, endIdx int) {
	if startIdx > endIdx {
		return
	}

	mid := startIdx + (endIdx-startIdx)/2
	if mid == 0 && nums[mid] > nums[mid+1] {
		res = append(res, mid)
		return
	}

	if mid == len(nums)-1 && nums[mid] > nums[mid-1] {
		res = append(res, mid)
		return
	}

	if mid > 0 && mid < len(nums)-1 && nums[mid] > nums[mid-1] &&
		nums[mid] > nums[mid+1] {
		res = append(res, mid)
		return
	}
	find(nums, startIdx, mid-1)
	find(nums, mid+1, endIdx)
}

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}

	res = res[:0]

	startIdx, endIdx := 0, len(nums)-1

	mid := startIdx + (endIdx-startIdx)/2

	find(nums, startIdx, mid-1)
	find(nums, mid, endIdx)

	if len(res) == 0 {
		return 0
	}

	return res[0]
}
