package main

func twoSum(numbers []int, target int) []int {
	var numMap = map[int]int{}
	for i := 0; i < len(numbers); i++ {
		numMap[numbers[i]] = i
	}

	for i := 0; i < len(numbers); i++ {
		if numMap[target-numbers[i]] > 0 && numMap[target-numbers[i]] != i {
			return []int{i + 1, numMap[target-numbers[i]] + 1}
		}
	}
	return []int{}
}
