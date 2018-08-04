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
func traverse(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	traverse(node.Left, res)
	*res = append(*res, node.Val)
	traverse(node.Right, res)
}
func minDiffInBST(root *TreeNode) int {
	traverseResult := []int{}
	traverse(root, &traverseResult)
	min := 300
	for i := 1; i < len(traverseResult); i++ {
		if traverseResult[i]-traverseResult[i-1] < min {
			min = traverseResult[i] - traverseResult[i-1]
		}
	}
	return min
}

func main() {

}
