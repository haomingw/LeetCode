package pb0009

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		args     args
		expected bool
	}{
		{
			args: args{
				x: 121,
			},
			expected: true,
		},
		{
			args: args{
				x: -121,
			},
			expected: false,
		},
		{
			args: args{
				x: 10,
			},
			expected: false,
		},
		{
			args: args{
				x: -101,
			},
			expected: false,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.x)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, isPalindrome(tc.args.x))
		})
	}
}
