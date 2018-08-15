package main

import "fmt"

var arr1 = map[int]bool{
	0: true,
	1: true,
	8: true,
	2: true,
	5: true,
	6: true,
	9: true,
}

var arr2 = map[int]bool{
	0: true,
	1: true,
	8: true,
}

func isSubSet(arr []int, t map[int]bool) bool {
	res := true
	for _, num := range arr {
		if t[num] == false {
			res = false
			break
		}
	}
	return res
}

func splitToInts(N int) []int {
	nStr := fmt.Sprint(N)
	res := []int{}
	for _, c := range nStr {
		res = append(res, int(c)-int('0'))
	}
	return res
}

func isGood(N int) bool {
	nums := splitToInts(N)
	//fmt.Println(nums)
	if isSubSet(nums, arr1) && !isSubSet(nums, arr2) {
		return true
	}

	return false
}

func rotatedDigits(N int) int {
	cnt := 0
	for i := 1; i <= N; i++ {
		if isGood(i) {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(rotatedDigits(10))
}
