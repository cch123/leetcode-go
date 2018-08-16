package main

import "fmt"

func main() {
	fmt.Println(isBipartite([][]int{[]int{1, 3}, []int{0, 2}, []int{1, 3}, []int{0, 2}}))
	fmt.Println(isBipartite([][]int{[]int{1, 2, 3}, []int{0, 2}, []int{0, 1, 3}, []int{0, 2}}))
	fmt.Println(isBipartite([][]int{[]int{},
		[]int{2, 4, 6}, []int{1, 4, 8, 9},
		[]int{7, 8}, []int{1, 2, 8, 9},
		[]int{6, 9}, []int{1, 5, 7, 8, 9},
		[]int{3, 6, 9}, []int{2, 3, 4, 6, 9},
		[]int{2, 4, 5, 6, 7, 8},
	}))
}

func isBipartite(graph [][]int) bool {
	bigraphLeft := map[int]bool{}
	bigraphRight := map[int]bool{}
	traversed := map[int]bool{}
	var useLeft bool
	for i := 0; i < len(graph); i++ {
		if traversed[i] == true {
			continue
		}
		// bfs this node and the connected dots
		useLeft, bigraphLeft[i] = false, true
		current := []int{i}
		for {
			next := []int{}
			allTraversed := true
			for j := 0; j < len(current); j++ {
				// connected to current[j]
				// traversed 表示该节点的子节点是否被访问过了
				if traversed[current[j]] != true {
					next = append(next, graph[current[j]]...)
					traversed[current[j]] = true
					allTraversed = false
				}
			}

			for j := 0; j < len(next); j++ {
				if useLeft {
					bigraphLeft[next[j]] = true
				} else {
					bigraphRight[next[j]] = true
				}

				if bigraphLeft[next[j]] == true && bigraphRight[next[j]] == true {
					//println(next[j])
					return false
				}
			}
			useLeft = !useLeft

			// 该连通分量已全部遍历到
			if allTraversed {
				break
			}

			current = next
			//fmt.Println(bigraphLeft, bigraphRight)
		}
	}
	return true
}
