package main

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	currentCnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			currentCnt++
			if currentCnt > max {
				max = currentCnt
			}
		} else {
			currentCnt = 0
		}
	}
	return max
}
