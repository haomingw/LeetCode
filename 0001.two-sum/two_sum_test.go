package pb0001

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args     args
		expected []int
	}{
		{
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
		{
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: []int{1, 2},
		},
		{
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			expected: []int{0, 1},
		},
		{
			args: args{
				nums:   []int{},
				target: 0,
			},
			expected: nil,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v,%v", tc.args.nums, tc.args.target)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, twoSum(tc.args.nums, tc.args.target))
		})
	}
}
