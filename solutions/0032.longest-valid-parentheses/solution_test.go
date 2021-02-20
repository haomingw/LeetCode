package pb0032

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args     args
		expected int
	}{
		{
			args: args{
				s: "(()",
			},
			expected: 2,
		},
		{
			args: args{
				s: ")()())",
			},
			expected: 4,
		},
		{
			args: args{
				s: "",
			},
			expected: 0,
		},
		{
			args: args{
				s: "()(()",
			},
			expected: 2,
		},
		{
			args: args{
				s: ")()())",
			},
			expected: 4,
		},
		{
			args: args{
				s: "(()(((()",
			},
			expected: 2,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.s)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, longestValidParentheses(tc.args.s))
		})
	}
}
