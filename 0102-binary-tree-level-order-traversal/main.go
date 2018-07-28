package main

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
type node struct {
	Val   int
	level int
}

func tranv(input []*TreeNode) []*TreeNode {
	res := []*TreeNode{}
	for i := 0; i < len(input); i++ {
		if input[i].Left != nil {
			res = append(res, input[i].Left)
		}
		if input[i].Right != nil {
			res = append(res, input[i].Right)
		}
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	var res = [][]int{}
	if root == nil {
		return res
	}

	level := 0
	currentLevelNodeList := []*TreeNode{root}
	for len(currentLevelNodeList) > 0 {
		// append to result
		currentLevelValList := []int{}
		for _, n := range currentLevelNodeList {
			currentLevelValList = append(currentLevelValList, n.Val)
		}
		res = append(res, currentLevelValList)

		// next level
		currentLevelNodeList = tranv(currentLevelNodeList)
		level++
	}
	return res
}
