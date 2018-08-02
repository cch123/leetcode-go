package main

import (
	"fmt"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	nStr := []byte(fmt.Sprint(n))
	res := -1
	has := false
	for i := 1; i < len(nStr); i++ {
		if nStr[i] > nStr[i-1] {
			has = true
		}
	}

	if has == false {
		return res
	}

	for i := len(nStr) - 1; i >= 1; i-- {
		if nStr[i] > nStr[i-1] {
			// 从 i 开始，向后，找一个最小的，且大于 i-1 元素的 index，交换
			minMax, minMaxIdx := nStr[i], i
			for j := i; j < len(nStr); j++ {
				if nStr[j] < minMax && nStr[j] > nStr[i-1] {
					minMax, minMaxIdx = nStr[j], j
				}
			}
			nStr[minMaxIdx], nStr[i-1] = nStr[i-1], nStr[minMaxIdx]
			// i 和其之后的都需要再次排序
			subSlice := nStr[i:]
			sort.Slice(subSlice, func(x, y int) bool {
				return subSlice[x] < subSlice[y]
			})
			break
		}
	}

	res64, _ := strconv.ParseInt(string(nStr), 10, 64)

	if res64 > 2147483648 {
		return -1
	}

	return int(res64)

}
