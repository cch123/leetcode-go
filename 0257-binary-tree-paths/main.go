package main

import "fmt"

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

func traverse(node *TreeNode, parentPath string, resultMap map[string]struct{}) {
	currentPath := parentPath + "->" + fmt.Sprint(node.Val)
	if parentPath == "" {
		currentPath = fmt.Sprint(node.Val)
	}
	if node.Left == nil && node.Right == nil {
		resultMap[currentPath] = struct{}{}
	}

	if node.Left != nil {
		traverse(node.Left, currentPath, resultMap)
	}

	if node.Right != nil {
		traverse(node.Right, currentPath, resultMap)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	var resultMap = map[string]struct{}{}
	var res = []string{}
	if root == nil {
		return res
	}

	traverse(root, "", resultMap)
	for k := range resultMap {
		res = append(res, k)
	}
	return res
}
