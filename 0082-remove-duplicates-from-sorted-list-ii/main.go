package main

func main() {
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var valMap = map[int]int{}
	cur := head
	for cur != nil {
		valMap[cur.Val]++
		cur = cur.Next
	}

	cur = head
	// change head
	for cur != nil && valMap[cur.Val] > 1 {
		cur = cur.Next
	}

	head = cur
	if head == nil {
		return head
	}

	cur = head.Next
	lastValid := head
	for cur != nil {
		if valMap[cur.Val] > 1 {
			lastValid.Next = cur.Next
			cur = cur.Next
			continue
		}

		lastValid = cur
		cur = cur.Next
	}

	return head
}
