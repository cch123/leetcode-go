package main

import "fmt"

func main() {
	fmt.Println(totalFruit([]int{1}))
	fmt.Println(totalFruit([]int{1, 2, 1}))
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}

func totalFruit(tree []int) int {
	if len(tree) <= 2 {
		return len(tree)
	}

	lastPosA, lastPosB := -1, -1
	curA, curB := -1, -1
	currentCnt, maxCnt := 0, 0
	for i := 0; i < len(tree); i++ {
		if tree[i] != curA && tree[i] != curB {
			// 选那个离当前位置最选的扔掉
			// 1, 2, 1, 1, 3 => 扔掉 2
			if lastPosA < lastPosB {
				currentCnt = i - lastPosA
				lastPosA = i
				curA = tree[i]
			} else {
				currentCnt = i - lastPosB
				lastPosB = i
				curB = tree[i]
			}
		} else if tree[i] == curA {
			lastPosA = i
			currentCnt++
		} else { // tree[i] == curB
			lastPosB = i
			currentCnt++
		}

		if currentCnt > maxCnt {
			maxCnt = currentCnt
		}
	}

	return maxCnt
}
