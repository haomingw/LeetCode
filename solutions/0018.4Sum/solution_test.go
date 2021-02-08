package pb0018

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args     args
		expected [][]int
	}{
		{
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			args: args{
				nums:   []int{0, 0, 0, 0},
				target: 0,
			},
			expected: [][]int{{0, 0, 0, 0}},
		},
		{
			args: args{
				nums:   []int{-2, -1, -1, 1, 1, 2, 2},
				target: 0,
			},
			expected: [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}},
		},
		{
			args: args{
				nums:   []int{},
				target: 0,
			},
			expected: [][]int{},
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.nums, tc.args.target)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, fourSum(tc.args.nums, tc.args.target))
		})
	}
}
