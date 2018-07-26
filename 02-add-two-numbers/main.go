package main

// ListNode ...
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{}
	cur := res
	for {
		if l1 != nil {
			cur.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			cur.Val += l2.Val
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			break
		}

		cur.Next = &ListNode{}
		cur = cur.Next
	}

	// 进位
	cur = res
	for cur != nil {
		if cur.Val >= 10 {
			if cur.Next == nil {
				cur.Next = &ListNode{}
			}

			cur.Val = cur.Val - 10
			cur.Next.Val++
		}

		cur = cur.Next
	}

	return res
}
