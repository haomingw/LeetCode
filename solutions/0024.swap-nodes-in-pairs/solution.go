package pb0024

import "github.com/haomingw/LeetCode/structures"

// ListNode alias
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	cur := dummy
	for cur != nil && cur.Next != nil {
		cur.Next = swap(cur.Next)
		cur = cur.Next.Next
	}
	return dummy.Next
}

func swap(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = head.Next.Next
	next.Next = head
	return next
}
