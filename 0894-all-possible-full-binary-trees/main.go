package main

func main() {}

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
func allPossibleFBT(N int) []*TreeNode {
	if N%2 == 0 {
		return []*TreeNode{}
	}

	if N == 1 {
		return []*TreeNode{
			&TreeNode{},
		}
	}

	if N == 3 {
		return []*TreeNode{
			&TreeNode{Left: &TreeNode{}, Right: &TreeNode{}},
		}
	}

	res := []*TreeNode{}
	for i := 1; i < N-1; i += 2 {
		lPossible := allPossibleFBT(i)
		rPossible := allPossibleFBT(N - 1 - i)
		for j := 0; j < len(lPossible); j++ {
			for k := 0; k < len(rPossible); k++ {
				cur := &TreeNode{
					Left:  lPossible[j],
					Right: rPossible[k],
				}
				res = append(res, cur)
			}
		}
	}
	return res
}
