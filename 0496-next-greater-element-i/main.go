package main

import "fmt"

func nextGreaterElement(findNums []int, nums []int) []int {
	var elemBiggerMap = map[int]int{}
	for i := 0; i < len(nums); i++ {
		biggerIdx := -1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				biggerIdx = nums[j]
				break
			}
		}
		elemBiggerMap[nums[i]] = biggerIdx
	}
	var res []int
	for i := 0; i < len(findNums); i++ {
		res = append(res, elemBiggerMap[findNums[i]])
	}
	return res
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
