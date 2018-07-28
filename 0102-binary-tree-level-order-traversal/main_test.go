package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	root     *TreeNode
	expected [][]int
}

func TestLevelOrder(t *testing.T) {
	cases := []TestCase{
		TestCase{
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: [][]int{
				[]int{3}, []int{9, 20}, []int{15, 7},
			},
		},
		TestCase{
			root:     nil,
			expected: [][]int{},
		},
	}

	for _, cas := range cases {
		res := levelOrder(cas.root)
		assert.Equal(t, cas.expected, res)
	}
}
