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
func traverse(node *TreeNode, sum *int) {
	if node == nil {
		return
	}

	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		*sum += node.Left.Val
	}
	traverse(node.Left, sum)
	traverse(node.Right, sum)
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	traverse(root, &sum)
	return sum
}
