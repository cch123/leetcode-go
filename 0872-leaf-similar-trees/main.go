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
var tree1LeafArr = []int{}
var tree2LeafArr = []int{}

func preTraversal(node *TreeNode, which int) {
	if node == nil {
		return
	}
	preTraversal(node.Left, which)
	preTraversal(node.Right, which)

	// is leaf
	if node.Left == nil && node.Right == nil {
		if which == 1 {
			tree1LeafArr = append(tree1LeafArr, node.Val)
		} else {
			tree2LeafArr = append(tree2LeafArr, node.Val)
		}
	}

}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	tree1LeafArr = tree1LeafArr[:0]
	tree1LeafArr = tree1LeafArr[:0]
	preTraversal(root1, 1)
	preTraversal(root2, 2)
	if len(tree1LeafArr) != len(tree2LeafArr) {
		return false
	}

	for i := 0; i < len(tree1LeafArr); i++ {
		if tree1LeafArr[i] != tree2LeafArr[i] {
			return false
		}
	}
	return true
}
