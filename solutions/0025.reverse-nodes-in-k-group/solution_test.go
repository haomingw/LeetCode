package pb0025

import (
	"fmt"
	"testing"

	"github.com/haomingw/LeetCode/structures"
	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		args     args
		expected *ListNode
	}{
		{
			args: args{
				head: structures.Array2Node([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			expected: structures.Array2Node([]int{2, 1, 4, 3, 5}),
		},
		{
			args: args{
				head: structures.Array2Node([]int{1, 2, 3, 4, 5}),
				k:    3,
			},
			expected: structures.Array2Node([]int{3, 2, 1, 4, 5}),
		},
		{
			args: args{
				head: structures.Array2Node([]int{1, 2, 3, 4, 5}),
				k:    1,
			},
			expected: structures.Array2Node([]int{1, 2, 3, 4, 5}),
		},
		{
			args: args{
				head: structures.Array2Node([]int{1}),
				k:    1,
			},
			expected: structures.Array2Node([]int{1}),
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.head)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, reverseKGroup(tc.args.head, tc.args.k))
		})
	}
}
