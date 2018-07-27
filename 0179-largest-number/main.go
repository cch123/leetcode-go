package main

import (
	"fmt"
	"sort"
)

func compare(x, y int) bool {
	iStr := fmt.Sprint(x)
	jStr := fmt.Sprint(y)
	if iStr+jStr > jStr+iStr {
		return true
	}
	return false
}

func largestNumber(nums []int) string {
	res := ""
	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i], nums[j])
	})

	for i := 0; i < len(nums); i++ {
		res += fmt.Sprint(nums[i])
	}

	// check whether all zero
	var allZeroFlag = true
	for i := 0; i < len(res); i++ {
		if res[i] != '0' {
			allZeroFlag = false
		}
	}

	if allZeroFlag {
		res = "0"
	}

	return res
}
