package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func dfs(parentNums []int, nums []int, res *[][]int) {
	for i := 0; i < len(nums); i++ {
		current := make([]int, len(parentNums)+1)
		copy(current, parentNums)
		current[len(current)-1] = nums[i]
		*res = append(*res, current)
		if i < len(nums) {
			dfs(current, nums[i+1:], res)
		}
	}
}

func subsets(nums []int) [][]int {
	res := [][]int{[]int{}}
	dfs([]int{}, nums, &res)
	return res
}
