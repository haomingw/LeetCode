package pb0005

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args     args
		expected []string
	}{
		{
			args: args{
				s: "babad",
			},
			expected: []string{"aba", "bab"},
		},
		{
			args: args{
				s: "cbbd",
			},
			expected: []string{"bb"},
		},
		{
			args: args{
				s: "a",
			},
			expected: []string{"a"},
		},
		{
			args: args{
				s: "ac",
			},
			expected: []string{"a", "c"},
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.s)
		t.Run(testname, func(t *testing.T) {
			assert.Contains(t, tc.expected, longestPalindrome(tc.args.s))
		})
		t.Run(testname, func(t *testing.T) {
			assert.Contains(t, tc.expected, longestPalindrome2(tc.args.s))
		})
	}
}
