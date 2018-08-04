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
func traverse(node *TreeNode, parentSum int, parentPath []int, target int, result *[][]int) {
	if node == nil {
		return
	}

	currentPath := append([]int{}, parentPath...)
	currentPath = append(currentPath, node.Val)
	currentSum := parentSum + node.Val

	// is leaf
	if node.Left == nil && node.Right == nil && currentSum == target {
		*result = append(*result, currentPath)
		return
	}

	if node.Left != nil {
		traverse(node.Left, currentSum, currentPath, target, result)
	}

	if node.Right != nil {
		traverse(node.Right, currentSum, currentPath, target, result)
	}
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	traverse(root, 0, []int{}, sum, &res)
	return res
}
