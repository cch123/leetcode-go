package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {

	sort.Slice(people, func(i, j int) bool {
		return people[i] > people[j]
	})

	// 从胖子开始，如果可以搭配当前最瘦的
	// 那就搭，搭不了就滚
	var totalBoat = 0
	var maxIdx = len(people) - 1
	var minIdx = 0
	for minIdx < maxIdx {
		totalBoat++
		if people[maxIdx]+people[minIdx] <= limit {
			minIdx++
			maxIdx--
			continue
		}
		minIdx++
	}
	// 说明还有一个哥们儿没带走
	if minIdx == maxIdx {
		totalBoat++
	}

	return totalBoat
}

func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}
