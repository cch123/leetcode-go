package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	l := countLength(head)
	midIdx := l / 2
	currentIdx := 0
	for {
		if currentIdx == midIdx {
			return head
		}
		head = head.Next
		currentIdx++
	}
}

func countLength(head *ListNode) int {
	counter := 0
	for head != nil {
		head = head.Next
		counter++
	}
	return counter
}
