package main

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	subtreeWithAllDeepest(tree)
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(n *TreeNode, parentPath []*TreeNode, leafMaxLen *int, leafNodeToPath map[*TreeNode][]*TreeNode) {
	if n == nil {
		return
	}

	currentPath := append([]*TreeNode{}, parentPath...)
	currentPath = append(currentPath, n)
	if n.Left == nil && n.Right == nil {
		currentLen := len(parentPath) + 1
		leafNodeToPath[n] = currentPath
		if currentLen > *leafMaxLen {
			*leafMaxLen = currentLen
		}
		return
	}
	traverse(n.Left, currentPath, leafMaxLen, leafNodeToPath)
	traverse(n.Right, currentPath, leafMaxLen, leafNodeToPath)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var leafNodeToPath = map[*TreeNode][]*TreeNode{}
	var leafMaxLen int
	traverse(root, []*TreeNode{}, &leafMaxLen, leafNodeToPath)

	previousPath := []*TreeNode{}
	//fmt.Println(leafNodeToPath)

	for _, path := range leafNodeToPath {
		if len(path) == leafMaxLen {
			if len(previousPath) == 0 {
				previousPath = path
				continue
			}
			// previousPath 和 path 的交集
			newPath := []*TreeNode{}
			x, y := 0, 0
			for x < len(previousPath) && y < len(path) && previousPath[x] == path[y] {
				newPath = append(newPath, previousPath[x])
				x++
				y++
			}
			previousPath = newPath
		}
	}
	//fmt.Println(previousPath)
	return previousPath[len(previousPath)-1]
}
