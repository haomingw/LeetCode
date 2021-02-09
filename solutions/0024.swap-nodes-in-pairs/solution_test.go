package pb0024

import (
	"fmt"
	"testing"

	"github.com/haomingw/LeetCode/structures"
	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		args     args
		expected *ListNode
	}{
		{
			args: args{
				head: structures.Array2Node([]int{1, 2, 3, 4}),
			},
			expected: structures.Array2Node([]int{2, 1, 4, 3}),
		},
		{
			args: args{
				head: structures.Array2Node([]int{}),
			},
			expected: structures.Array2Node([]int{}),
		},
		{
			args: args{
				head: structures.Array2Node([]int{1}),
			},
			expected: structures.Array2Node([]int{1}),
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.head)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, swapPairs(tc.args.head))
		})
	}
}
