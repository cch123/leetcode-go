package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 4,
		//Left: &TreeNode{Val: 6},
		Right: &TreeNode{Val: 4},
	}
	fmt.Println(tree2str(tree))
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

func traverse(n *TreeNode) string {
	if n == nil {
		return ""
	}

	if n.Left == nil && n.Right == nil {
		return fmt.Sprint(n.Val)
	}

	str := fmt.Sprint(n.Val) + "(" + traverse(n.Left) + ")"
	if n.Right != nil {
		str += "(" + traverse(n.Right) + ")"
	}

	return str
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	return traverse(t)
}
