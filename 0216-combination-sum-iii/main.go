package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 2))
	fmt.Println(combinationSum3(2, 18))
}

func nextCombination(zoArr []int) bool {
	// 从左向右找第一个 10
	// swap 10 => 01
	// 把 01 之前所有的 1 移到开头位置
	found := false
	for i := 0; i < len(zoArr)-1; i++ {
		if zoArr[i] == 1 && zoArr[i+1] == 0 {
			if i >= 8 {
				break
			}
			zoArr[i], zoArr[i+1] = zoArr[i+1], zoArr[i]
			posForOne := 0
			for j := 0; j < i; j++ {
				if zoArr[j] == 1 {
					zoArr[j] = 0
					zoArr[posForOne] = 1
					posForOne++
				}
			}
			found = true
			break
		}
	}
	fmt.Println(zoArr)
	return found
}

func calcSum(zoArr []int) int {
	sum := 0
	for i := 0; i < len(zoArr); i++ {
		if zoArr[i] == 1 {
			sum += i + 1
		}
	}
	return sum
}

func combinationSum3(k int, n int) [][]int {
	if k >= n {
		return [][]int{}
	}
	zoArr := make([]int, n)
	// init
	for i := 0; i < k; i++ {
		zoArr[i] = 1
	}

	var res = [][]int{}
	for {
		if calcSum(zoArr) == n {
			currentArr := []int{}
			for i := 0; i < len(zoArr); i++ {
				if zoArr[i] == 1 {
					currentArr = append(currentArr, i+1)
				}
			}
			res = append(res, currentArr)
		}

		if !nextCombination(zoArr) {
			break
		}
	}
	return res
}
