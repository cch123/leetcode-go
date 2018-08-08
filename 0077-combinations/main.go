package main

import (
	"sort"
)

func main() {}
func getResult(pickArr []int) []int {
	currentCombi := []int{}
	for j := 0; j < len(pickArr); j++ {
		if pickArr[j] == 1 {
			currentCombi = append(currentCombi, j+1)
		}
	}
	return currentCombi
}

func combine(n int, k int) [][]int {
	pickArr := make([]int, n)
	for i := 0; i < k; i++ {
		pickArr[i] = 1
	}

	res := append([][]int{}, getResult(pickArr))

	for {
		// find the first 10
		// swap 10 -> 01
		// move all 1 left of 01 to the left
		found := false
		for i := 0; i < len(pickArr)-1; i++ {
			if pickArr[i] == 1 && pickArr[i+1] == 0 {
				found = true
				pickArr[i], pickArr[i+1] = 0, 1
				onePos := 0
				for j := 0; j < i; j++ {
					if pickArr[j] == 1 {
						pickArr[j], pickArr[onePos] = pickArr[onePos], pickArr[j]
						onePos++
					}
				}
				sort.Slice(pickArr[:i], func(x, y int) bool {
					return pickArr[x] > pickArr[y]
				})
				break
			}
		}

		if !found {
			break
		}
		// found
		res = append(res, getResult(pickArr))
	}
	return res
}
