package pb0002

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	current := dummy
	a, b, carry := 0, 0, 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			a = 0
		} else {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			b = 0
		} else {
			b = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (a + b + carry) % 10}
		current = current.Next
		carry = (a + b + carry) / 10
	}
	return dummy.Next
}
