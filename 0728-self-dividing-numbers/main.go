package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	res := []int{}
outer:
	for i := left; i <= right; i++ {
		iStr := fmt.Sprint(i)
		for j := 0; j < len(iStr); j++ {
			if iStr[j] == '0' {
				continue outer
			}

			if i%(int(iStr[j])-int('0')) != 0 {
				continue outer
			}
		}
		res = append(res, i)
	}
	return res
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
