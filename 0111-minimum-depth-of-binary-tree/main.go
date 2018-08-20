package main

func main() {}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(n *TreeNode, parentDepth int, minDepth *int) {
	if n == nil {
		return
	}
	currentDep := parentDepth + 1
	if n.Left == nil && n.Right == nil {
		if currentDep < *minDepth {
			*minDepth = currentDep
		}
		return
	}

	traverse(n.Left, currentDep, minDepth)
	traverse(n.Right, currentDep, minDepth)
}

func minDepth(root *TreeNode) int {
	var res = 999999999
	traverse(root, 0, &res)
	return res
}
