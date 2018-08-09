package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1}, 0))
}

func subarraySum(nums []int, k int) int {
	sumMap := map[int]int{0: 1}
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		fmt.Println(sum, k, sumMap)
		if sumMap[sum-k] > 0 {
			res += sumMap[sum-k]
		}
		sumMap[sum]++
	}
	return res
}
