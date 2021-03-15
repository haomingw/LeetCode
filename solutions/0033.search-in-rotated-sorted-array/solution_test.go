package pb0033

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestValidParentheses(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args     args
		expected int
	}{
		{
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			expected: 4,
		},
		{
			args: args{
				nums:   []int{5, 6, 7, 0, 1, 2, 4},
				target: 5,
			},
			expected: 0,
		},
		{
			args: args{
				nums:   []int{5, 6, 7, 0, 1, 2, 4},
				target: 1,
			},
			expected: 4,
		},
		{
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			expected: -1,
		},
		{
			args: args{
				nums:   []int{2, 1},
				target: 1,
			},
			expected: 1,
		},
		{
			args: args{
				nums:   []int{1},
				target: 0,
			},
			expected: -1,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.nums, tc.args.target)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, search(tc.args.nums, tc.args.target))
		})
	}
}
