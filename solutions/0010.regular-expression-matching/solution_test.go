package pb0010

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		args     args
		expected bool
	}{
		{
			args: args{
				s: "aa",
				p: "a",
			},
			expected: false,
		},
		{
			args: args{
				s: "aa",
				p: "a*",
			},
			expected: true,
		},
		{
			args: args{
				s: "ab",
				p: ".*",
			},
			expected: true,
		},
		{
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			expected: true,
		},
		{
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			expected: false,
		},
		{
			args: args{
				s: "mississippi",
				p: "mis*is*ip*.",
			},
			expected: true,
		},
		{
			args: args{
				s: "aaa",
				p: "ab*a*c*a",
			},
			expected: true,
		},
		{
			args: args{
				s: "aaa",
				p: "ab*ac*a",
			},
			expected: true,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.s, tc.args.p)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, isMatch(tc.args.s, tc.args.p))
		})
	}
}
