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
func isValidBSTWithBounds(root *TreeNode, lowerBound int, upperBound int) bool {
	if root == nil {
		return true
	}

	if (lowerBound != -1 && root.Val <= lowerBound) || (upperBound != -1 && root.Val >= upperBound) {
		return false
	}

	leftUpperBound := root.Val
	if upperBound != -1 && upperBound <= leftUpperBound {
		leftUpperBound = upperBound
	}

	rightLowerBound := root.Val
	if lowerBound != -1 && lowerBound >= rightLowerBound {
		rightLowerBound = lowerBound
	}

	return isValidBSTWithBounds(root.Left, lowerBound, leftUpperBound) &&
		isValidBSTWithBounds(root.Right, rightLowerBound, upperBound)

}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// left
	leftUpperBound := root.Val
	if root.Left != nil {
		if !isValidBSTWithBounds(root.Left, -1, leftUpperBound) {
			return false
		}
	}

	// right
	rightLowerBound := root.Val
	if root.Right != nil {
		if !isValidBSTWithBounds(root.Right, rightLowerBound, -1) {
			return false
		}
	}

	return true
}
