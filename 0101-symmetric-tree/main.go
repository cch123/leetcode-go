package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseTree(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil || root.Right != nil {
		root.Left, root.Right = root.Right, root.Left
	}

	reverseTree(root.Left)
	reverseTree(root.Right)
}

func traverseAndCompare(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil && node2 != nil {
		return false
	}

	if node1 != nil && node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	return traverseAndCompare(node1.Left, node2.Left) &&
		traverseAndCompare(node1.Right, node2.Right)

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	reverseTree(root.Right)
	return traverseAndCompare(root.Left, root.Right)
}
