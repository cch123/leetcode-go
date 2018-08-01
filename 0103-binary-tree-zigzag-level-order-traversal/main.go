package main

import "fmt"

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}

	currentLevelNodes := []*TreeNode{root}
	for len(currentLevelNodes) > 0 {
		traverseResult := []int{}
		var nextLevelNodes = []*TreeNode{}
		for i := 0; i < len(currentLevelNodes); i++ {
			traverseResult = append(traverseResult, currentLevelNodes[i].Val)
			if currentLevelNodes[i].Left != nil {
				nextLevelNodes = append(nextLevelNodes, currentLevelNodes[i].Left)
			}
			if currentLevelNodes[i].Right != nil {
				nextLevelNodes = append(nextLevelNodes, currentLevelNodes[i].Right)

			}
		}
		res = append(res, traverseResult)
		currentLevelNodes = nextLevelNodes
	}
	for i := 0; i < len(res); i++ {
		if i%2 == 0 {
			continue
		}

		for j := 0; j < len(res[i])/2; j++ {
			res[i][j], res[i][len(res[i])-1-j] = res[i][len(res[i])-1-j], res[i][j]
		}
	}
	return res
}

func main() {
	tree := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(zigzagLevelOrder(tree))
	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  1,
				Left: &TreeNode{},
			},
			Right: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{},
				},
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(zigzagLevelOrder(tree))
}
