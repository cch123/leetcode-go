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
func removeElements(head *ListNode, val int) *ListNode {

	// decide head
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

func main() {
	root := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 6,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5,
							Next: &ListNode{Val: 6}}}}}}}
	removeElements(root, 3)
	for root != nil {
		println(root.Val)
		root = root.Next
	}
}
