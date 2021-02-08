package pb0020

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args     args
		expected bool
	}{
		{
			args: args{
				s: "()",
			},
			expected: true,
		},
		{
			args: args{
				s: "()[]{}",
			},
			expected: true,
		},
		{
			args: args{
				s: "(]",
			},
			expected: false,
		},
		{
			args: args{
				s: "([)]",
			},
			expected: false,
		},
		{
			args: args{
				s: "{[]}",
			},
			expected: true,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.s)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, isValid(tc.args.s))
		})
	}
}
