package main

func twoSum(nums []int, target int) []int {
	var numMap = map[int][]int{}
	for idx, num := range nums {
		numMap[num] = append(numMap[num], idx)
	}

	var indexMap = []int{}

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		left := target - cur
		if cur == left {
			if len(numMap[cur]) != 2 {
				continue
			}
			indexMap = numMap[cur]
			break
		} else if _, ok := numMap[left]; ok {
			indexMap = []int{i, numMap[left][0]}
			break
		}
	}

	return indexMap
}
