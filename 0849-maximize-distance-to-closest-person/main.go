package main

import "fmt"

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))
}

func maxDistToClosest(seats []int) int {
	f := make([]int, len(seats))
	for i := 0; i < len(f); i++ {
		f[i] = len(seats) + 10
	}
	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			continue
		}
		f[i] = 0

		// seats[i] == 0
		// 向左
		for j := i - 1; j >= 0; j-- {
			if seats[j] == 1 {
				break
			}
			if f[j] > i-j {
				f[j] = i - j
			} else {
				break
			}
		}
		// 向右
		for j := i + 1; j < len(seats); j++ {
			if seats[j] == 1 {
				i = j - 1
				break
			}
			f[j] = j - i
		}
	}
	max := 0
	for i := 0; i < len(f); i++ {
		if f[i] > max {
			max = f[i]
		}
	}
	return max
}
