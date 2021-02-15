package pb0031

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args     args
		expected []int
	}{
		{
			args: args{
				nums: []int{1, 2, 3},
			},
			expected: []int{1, 3, 2},
		},
		{
			args: args{
				nums: []int{3, 2, 1},
			},
			expected: []int{1, 2, 3},
		},
		{
			args: args{
				nums: []int{1},
			},
			expected: []int{1},
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.nums)
		t.Run(testname, func(t *testing.T) {
			nextPermutation(tc.args.nums)
			assert.Equal(t, tc.expected, tc.args.nums)
		})
	}
}
