package pb0015

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args     args
		expected [][]int
	}{
		{
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			args: args{
				nums: []int{-1, 0, 1, 1, 2, -1, -1, -4},
			},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			args: args{
				nums: []int{},
			},
			expected: [][]int{},
		},
		{
			args: args{
				nums: []int{0},
			},
			expected: [][]int{},
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.nums)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, threeSum(tc.args.nums))
		})
	}
}
