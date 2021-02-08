package pb0019

import (
	"fmt"
	"testing"

	"github.com/haomingw/LeetCode/structures"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		args     args
		expected *ListNode
	}{
		{
			args: args{
				head: structures.Array2Node([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			expected: structures.Array2Node([]int{1, 2, 3, 5}),
		},
		{
			args: args{
				head: structures.Array2Node([]int{1}),
				n:    1,
			},
			expected: structures.Array2Node([]int{}),
		},
		{
			args: args{
				head: structures.Array2Node([]int{1, 2}),
				n:    1,
			},
			expected: structures.Array2Node([]int{1}),
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.head, tc.args.n)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, removeNthFromEnd(tc.args.head, tc.args.n))
		})
	}
}
