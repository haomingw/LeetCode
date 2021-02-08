package pb0019

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	prev, slow, fast := dummy, head, head
	for i := 0; i < n-1; i++ {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}
	prev.Next = slow.Next
	return dummy.Next
}
