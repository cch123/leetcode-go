package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	result := addTwoNumbers(l1, l2)
	assert.Equal(t, result.Val, 7)
	assert.Equal(t, result.Next.Val, 0)
	assert.Equal(t, result.Next.Next.Val, 8)

	l1 = &ListNode{
		Val:  0,
		Next: nil,
	}

	l2 = &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 3,
		},
	}

	result = addTwoNumbers(l1, l2)
	assert.Equal(t, result.Val, 7)
	assert.Equal(t, result.Next.Val, 3)

}
