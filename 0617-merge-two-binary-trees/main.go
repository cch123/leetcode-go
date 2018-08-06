package main

func main() {

}

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
func traverse(n1 *TreeNode, n2 *TreeNode) {
	// TODO, fix this?
	if n1 == nil || n2 == nil {
		return
	}
	n1.Val += n2.Val

	if n1.Left != nil && n2.Left != nil {
		traverse(n1.Left, n2.Left)
	}

	/*
		if n1.Left == nil && n2.Left == nil {
			// do nothing
		}

		if n1.Left != nil && n2.Left == nil {
			// do nothing
		}
	*/

	if n1.Left == nil && n2.Left != nil {
		n1.Left = n2.Left
		n2.Left = nil
	}

	if n1.Right != nil && n2.Right != nil {
		traverse(n1.Right, n2.Right)
	}

	if n1.Right == nil && n2.Right != nil {
		n1.Right = n2.Right
		n2.Right = nil
	}

	/*
		if n1.Right != nil && n2.Right == nil {
			// do nothing
		}

		if n1.Right == nil && n2.Right == nil {
			// do nothing
		}
	*/
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		// 让 t1 尽量不是 nil
		t1, t2 = t2, t1
	}
	traverse(t1, t2)
	return t1
}
