package main

import "fmt"

type lrSum struct {
	LSum int
	RSum int
}

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	sumArr := make([]lrSum, len(nums))

	rSum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		rSum += nums[i]
		sumArr[i].RSum = rSum
	}

	lSum := 0
	for i := 0; i < len(nums); i++ {
		lSum += nums[i]
		sumArr[i].LSum = lSum
	}
	fmt.Println(sumArr)

	if sumArr[1].RSum == 0 {
		return 0
	}

	for i := 1; i < len(nums)-1; i++ {
		if sumArr[i-1].LSum == sumArr[i+1].RSum {
			return i
		}
	}

	if sumArr[len(sumArr)-2].LSum == 0 {
		return len(sumArr) - 1
	}

	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{-1, -1, -1, 0, 1, 1}))
	fmt.Println(pivotIndex([]int{-1, -1, -1, 1, 1, 1}))
}
