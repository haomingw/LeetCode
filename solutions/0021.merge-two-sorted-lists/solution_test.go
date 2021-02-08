package pb0021

import (
	"fmt"
	"testing"

	"github.com/haomingw/LeetCode/structures"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		args     args
		expected *ListNode
	}{
		{
			args: args{
				l1: structures.Array2Node([]int{1, 2, 4}),
				l2: structures.Array2Node([]int{1, 3, 4}),
			},
			expected: structures.Array2Node([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			args: args{
				l1: structures.Array2Node([]int{}),
				l2: structures.Array2Node([]int{}),
			},
			expected: structures.Array2Node([]int{}),
		},
		{
			args: args{
				l1: structures.Array2Node([]int{}),
				l2: structures.Array2Node([]int{0}),
			},
			expected: structures.Array2Node([]int{0}),
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.l1, tc.args.l2)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, mergeTwoLists(tc.args.l1, tc.args.l2))
		})
	}
}
