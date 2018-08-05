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
func searchBST(root *TreeNode, val int) *TreeNode {
	for {
		if root == nil {
			return nil
		}

		if val == root.Val {
			return root
		}
		if val < root.Val {
			root = root.Left
			continue
		}

		if val > root.Val {
			root = root.Right
			continue
		}
	}
}
