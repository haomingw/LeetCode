package structures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray2Node(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args     args
		expected *ListNode
	}{
		{
			args: args{
				nums: []int{},
			},
			expected: nil,
		},
		{
			args: args{
				nums: []int{1, 2, 3},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	equal := func(a *ListNode, b *ListNode) bool {
		for a != nil && b != nil {
			if a.Val != b.Val {
				return false
			}
			a = a.Next
			b = b.Next
		}
		return a == nil && b == nil
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.nums)
		t.Run(testname, func(t *testing.T) {
			assert.True(t, equal(tc.expected, Array2Node(tc.args.nums)))
		})
	}
}

func TestNode2Array(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		args     args
		expected []int
	}{
		{
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			expected: []int{1, 2, 3},
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.expected)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, Node2Array(tc.args.head))
		})
	}
	tooLong := Array2Node(make([]int, 11))
	assert.Panics(t, func() { Node2Array(tooLong) }, "Current depth 11 exceeds limit 10")
}

func TestListNodeRepr(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		args     args
		expected string
	}{
		{
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			expected: "[1,2,3]",
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.expected)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, fmt.Sprint(tc.args.head))
		})
	}
}
