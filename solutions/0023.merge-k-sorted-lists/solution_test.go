package pb0023

import (
	"fmt"
	"testing"

	"github.com/haomingw/LeetCode/structures"
	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		args     args
		expected *ListNode
	}{
		{
			args: args{
				lists: []*ListNode{
					structures.Array2Node([]int{1, 4, 5}),
					structures.Array2Node([]int{1, 3, 4}),
					structures.Array2Node([]int{2, 6}),
				},
			},
			expected: structures.Array2Node([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			args: args{
				lists: []*ListNode{
					structures.Array2Node([]int{1, 4, 5}),
					structures.Array2Node([]int{1, 3, 6}),
					structures.Array2Node([]int{2, 4}),
				},
			},
			expected: structures.Array2Node([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			args: args{
				lists: []*ListNode{},
			},
			expected: nil,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.lists)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, mergeKLists(tc.args.lists))
		})
	}
}
