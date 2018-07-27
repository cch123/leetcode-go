package main

func maxSubArray(nums []int) int {
	sumMax := make([]int, len(nums))
	sumMaxResult := nums[0]
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sumMax[i] = nums[i]
			if sumMax[i] > sumMaxResult {
				sumMaxResult = sumMax[i]
			}
		} else {
			if sumMax[i-1] > 0 {
				sumMax[i] = sumMax[i-1] + nums[i]
			} else {
				sumMax[i] = nums[i]
			}
		}

		if sumMax[i] > sumMaxResult {
			sumMaxResult = sumMax[i]
		}
	}
	return sumMaxResult
}
