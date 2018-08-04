package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(node *TreeNode, parentSum int, target int) bool {
    if node == nil {
        return false
    }
    
    currentSum := parentSum + node.Val
    
    if node.Left == nil && node.Right == nil && currentSum == target {
        return true
    }
    
    return traverse(node.Left, currentSum, target) || traverse(node.Right, currentSum, target)
}

func hasPathSum(root *TreeNode, sum int) bool {
    return traverse(root, 0, sum)
}
