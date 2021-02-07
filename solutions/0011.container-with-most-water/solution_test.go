package pb0011

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		args     args
		expected int
	}{
		{
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			expected: 49,
		},
		{
			args: args{
				height: []int{1, 1},
			},
			expected: 1,
		},
		{
			args: args{
				height: []int{4, 3, 2, 1, 4},
			},
			expected: 16,
		},
		{
			args: args{
				height: []int{1, 2, 1},
			},
			expected: 2,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.height)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, maxArea(tc.args.height))
		})
	}
}
