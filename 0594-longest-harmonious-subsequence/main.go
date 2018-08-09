package main

func main() {}

func findLHS(nums []int) int {
	var valMap = map[int]int{}
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		valMap[num]++
		if valMap[num+1] == 0 && valMap[num-1] == 0 {
			continue
		}

		if valMap[num]+valMap[num-1] > maxLen {
			maxLen = valMap[num] + valMap[num-1]
		}

		if valMap[num]+valMap[num+1] > maxLen {
			maxLen = valMap[num] + valMap[num+1]
		}
	}
	return maxLen
}
