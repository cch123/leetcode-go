package main

import "fmt"

func main() {
	fmt.Println(possibleBipartition(4, [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 4}}))
}

func getNextLevel(dislikeMap map[int]map[int]bool, current map[int]bool) map[int]bool {
	nextLevel := map[int]bool{}
	// needToGoDeeper := false
	for elem := range current {
		dislikeForI := dislikeMap[elem]
		for k := range dislikeForI {
			nextLevel[k] = true
		}
	}
	return nextLevel
}

func hasElemWhichNotInMap(elems map[int]bool, valMap map[int]bool) bool {
	for elem := range elems {
		if valMap[elem] == false {
			return true
		}
	}

	return false
}

func possibleBipartition(N int, dislikes [][]int) bool {
	if len(dislikes) == 0 {
		return true
	}

	var dislikeMap = map[int]map[int]bool{}
	for i := 0; i < len(dislikes); i++ {
		x, y := dislikes[i][0], dislikes[i][1]
		if dislikeMap[x] == nil {
			dislikeMap[x] = make(map[int]bool)
		}
		if dislikeMap[y] == nil {
			dislikeMap[y] = make(map[int]bool)
		}
		dislikeMap[x][y] = true
		dislikeMap[y][x] = true
	}
	// fmt.Println(dislikeMap)

	mapA := map[int]bool{}
	mapB := map[int]bool{}
	for i := 1; i <= N; i++ {
		if mapA[i] == true || mapB[i] == true {
			continue
		}

		// 从 i 开始广度优先搜索
		currentLevel := map[int]bool{i: true}
		mapA[i] = true
		cnt := 1
		for {
			nextLevel := getNextLevel(dislikeMap, currentLevel)
			//fmt.Println("next level", nextLevel)
			var m map[int]bool
			var n map[int]bool
			if cnt == 0 {
				// 放 mapA
				if !hasElemWhichNotInMap(nextLevel, mapA) {
					break
				}
				m, n, cnt = mapA, mapB, 1
			} else {
				// 放 mapB
				if !hasElemWhichNotInMap(nextLevel, mapB) {
					break
				}
				m, n, cnt = mapB, mapA, 0
			}

			for elem := range nextLevel {
				m[elem] = true

				if n[elem] == true {
					return false
				}
			}
			currentLevel = nextLevel
			//fmt.Println(m, n)
		}
	}

	//	fmt.Println(mapA, mapB)

	return true
}
