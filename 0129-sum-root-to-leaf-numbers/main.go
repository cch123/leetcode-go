package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(node *TreeNode, parentSum int, resultArr *[]int) {
	if node == nil {
		return
	}

	currentSum := parentSum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		*resultArr = append(*resultArr, currentSum)
	}

	traverse(node.Left, currentSum, resultArr)
	traverse(node.Right, currentSum, resultArr)
}

func sumNumbers(root *TreeNode) int {
	leafNums := []int{}
	traverse(root, 0, &leafNums)
	sum := 0
	for _, num := range leafNums {
		sum += num
	}
	return sum
}
