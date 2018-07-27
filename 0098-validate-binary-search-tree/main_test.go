package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input    *TreeNode
	expected bool
}

func TestIsValidBST(t *testing.T) {
	cases := []TestCase{
		TestCase{
			input: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: true,
		},
		TestCase{
			input: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{Val: 4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
		},
	}

	for _, cas := range cases {
		res := isValidBST(cas.input)
		assert.Equal(t, cas.expected, res)
	}
}
