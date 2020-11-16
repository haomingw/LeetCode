package structures

import "fmt"

// ListNode is a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprint(Node2Array(l))
}

// Array2Node converts an array to a LinkedList
func Array2Node(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Node2Array converts a LinkedList to an array
func Node2Array(head *ListNode) []int {
	res := []int{}
	depth, limit := 0, 10
	for head != nil {
		depth++
		if depth > limit {
			msg := fmt.Sprintf("Current depth %d exceeds limit %d", depth, limit)
			panic(msg)
		}
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
