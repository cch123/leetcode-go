package main

import "fmt"

func main() {
	fmt.Println(swimInWater([][]int{[]int{0, 2}, []int{1, 3}}))
	fmt.Println(swimInWater([][]int{[]int{10, 12, 4, 6}, []int{9, 11, 3, 5}, []int{1, 7, 13, 8}, []int{2, 0, 15, 14}}))
}

type point struct {
	x, y int
}

type r struct {
	left, right, down, up bool
}

func swimInWater(grid [][]int) int {
	reached := map[point]bool{
		point{0, 0}: true,
	}
	traversed := map[point]bool{point{x: 0, y: 0}: true}
	res := grid[0][0]
	for {
		// reached 的所有 adjacent node
		// 均需要加入到 reached
		curReached := []point{}
		for p := range reached {
			curReached = append(curReached, p)
		}
		for {
			newReached := []point{}
			//fmt.Println(traversed)
			for _, p := range curReached {
				// 左
				leftP := point{x: p.x, y: p.y - 1}
				if !traversed[leftP] {
					if leftP.y >= 0 && !reached[leftP] && grid[leftP.x][leftP.y] <= res {
						newReached = append(newReached, leftP)
						traversed[leftP] = true
					}
				}
				// 下
				downP := point{x: p.x + 1, y: p.y}
				if !traversed[downP] {
					if downP.x < len(grid) && !reached[downP] && grid[downP.x][downP.y] <= res {
						newReached = append(newReached, downP)
						traversed[downP] = true
					}
				}
				// 上
				upP := point{x: p.x - 1, y: p.y}
				if !traversed[upP] {
					if upP.x >= 0 && !reached[upP] && grid[upP.x][upP.y] <= res {
						newReached = append(newReached, upP)
						traversed[upP] = true
					}
				}
				// 右
				rightP := point{x: p.x, y: p.y + 1}
				if !traversed[rightP] {
					if rightP.y < len(grid[0]) && !reached[rightP] && grid[rightP.x][rightP.y] <= res {
						newReached = append(newReached, rightP)
						traversed[rightP] = true
					}
				}
			}
			if len(newReached) == 0 {
				break
			}

			// 把 newReached 加到 reach map
			// 然后 curReached = newReached
			for _, p := range newReached {
				reached[p] = true
			}
			curReached = newReached
		}
		//fmt.Println(res, reached)
		if reached[point{x: len(grid) - 1, y: len(grid) - 1}] == true {
			return res
		}
		res++
	}
}
