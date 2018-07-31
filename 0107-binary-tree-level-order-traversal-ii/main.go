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
func levelOrderBottom(root *TreeNode) [][]int {
	var res = [][]int{}
	currentLevel := []*TreeNode{root}
	for len(currentLevel) > 0 {
		nextLevel := []*TreeNode{}
		currentLevelTraverseResult := []int{}

		for _, n := range currentLevel {
			if n != nil {
				currentLevelTraverseResult = append(currentLevelTraverseResult, n.Val)
			}

			if n != nil && n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}

			if n != nil && n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
		}

		if len(currentLevelTraverseResult) > 0 {
			res = append(res, currentLevelTraverseResult)
		}

		currentLevel = nextLevel
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
