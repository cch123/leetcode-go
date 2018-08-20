package main

func main() {}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func depth(n *TreeNode) int {
	if n == nil {
		return 0
	}

	leftDepth := depth(n.Left)
	rightDepth := depth(n.Right)
	if leftDepth == -1 || rightDepth == -1 || abs(leftDepth-rightDepth) > 1 {
		return -1
	}
	return max(leftDepth, rightDepth) + 1

}

func isBalanced(root *TreeNode) bool {
	return depth(root) != -1
}
