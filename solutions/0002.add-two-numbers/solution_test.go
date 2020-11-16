package pb0002

import (
	"fmt"
	"testing"

	"github.com/haomingw/LeetCode/structures"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
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
				l1: structures.Array2Node([]int{2, 4, 3}),
				l2: structures.Array2Node([]int{5, 6, 4}),
			},
			expected: structures.Array2Node([]int{7, 0, 8}),
		},
		{
			args: args{
				l1: structures.Array2Node([]int{1}),
				l2: structures.Array2Node([]int{9, 9, 9}),
			},
			expected: structures.Array2Node([]int{0, 0, 0, 1}),
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v,%v", tc.args.l1, tc.args.l2)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, addTwoNumbers(tc.args.l1, tc.args.l2))
		})
	}
}
