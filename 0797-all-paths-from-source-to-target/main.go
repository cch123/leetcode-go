package main

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{[]int{1, 2}, []int{3}, []int{3}, []int{}}))
}

func traverse(elem, target int, parentPath []int,
	traversed map[int]bool, res *[][]int, graph [][]int) {

	currentPath := append([]int{}, parentPath...)
	currentPath = append(currentPath, elem)

	if elem == target {
		resultPathElem := append([]int{}, currentPath...)
		*res = append(*res, resultPathElem)
	}

	children := graph[elem]
	for i := 0; i < len(children); i++ {
		traverse(children[i], target, currentPath, traversed, res, graph)
	}
}

func allPathsSourceTarget(graph [][]int) [][]int {
	if len(graph) == 0 {
		return [][]int{}
	}

	var res = [][]int{}
	target := len(graph) - 1
	traversed := map[int]bool{}
	traverse(0, target, []int{}, traversed, &res, graph)

	return res
}
