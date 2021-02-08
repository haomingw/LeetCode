package pb0016

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumClosest(t *testing.T) {
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
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			expected: 2,
		},
		{
			args: args{
				nums:   []int{0, 2, 1, -3},
				target: 1,
			},
			expected: 0,
		},
		{
			args: args{
				nums:   []int{0, 2, 1, -3, -3},
				target: 1,
			},
			expected: 0,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.nums, tc.args.target)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, threeSumClosest(tc.args.nums, tc.args.target))
		})
	}
}
