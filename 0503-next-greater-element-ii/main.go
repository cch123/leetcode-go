package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	numsMore := append([]int{}, nums...)
	numsMore = append(numsMore, nums...)
	result := []int{}
	for i := 0; i < len(nums); i++ {
		biggerResult := -1
		for j := i + 1; j < i+len(nums); j++ {
			if numsMore[j] > nums[i] {
				biggerResult = numsMore[j]
				break
			}
		}
		result = append(result, biggerResult)
	}
	return result
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}
