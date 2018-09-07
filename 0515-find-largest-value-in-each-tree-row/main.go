package main

func main() {}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	cur := []*TreeNode{root}
	next := []*TreeNode{}
	res := []int{}
	for {
		if len(cur) == 0 {
			break
		}
		max := cur[0].Val
		// find biggest of cur
		for i := 0; i < len(cur); i++ {
			if cur[i].Val > max {
				max = cur[i].Val
			}
			if cur[i].Left != nil {
				next = append(next, cur[i].Left)
			}
			if cur[i].Right != nil {
				next = append(next, cur[i].Right)
			}
		}
		res = append(res, max)
		cur = append([]*TreeNode{}, next...)
		next = next[:0]
	}
	return res
}
