package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
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

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{[]int{}}
	dfs([]int{}, nums, &res)
	sort.Slice(res, func(x, y int) bool {
		if len(res[x]) < len(res[y]) {
			return true
		}
		if len(res[x]) > len(res[y]) {
			return false
		}
		for i := 0; i < len(res[x]); i++ {
			if res[x][i] < res[y][i] {
				return true
			}
			if res[x][i] > res[y][i] {
				return false
			}
		}
		return false
	})

	finalRes := [][]int{res[0]}

	for i := 1; i < len(res); i++ {
		if equal(res[i], res[i-1]) {
			// skip
			continue
		}
		finalRes = append(finalRes, res[i])
	}
	return finalRes
}
