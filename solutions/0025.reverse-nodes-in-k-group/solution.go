package pb0025

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
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(head *ListNode, tail *ListNode) *ListNode {
	prev := tail
	for head != tail {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
