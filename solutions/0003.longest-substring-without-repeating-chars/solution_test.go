package pb0003

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args     args
		expected int
	}{
		{
			args: args{
				s: "abcabcbb",
			},
			expected: 3,
		},
		{
			args: args{
				s: "bbbbb",
			},
			expected: 1,
		},
		{
			args: args{
				s: "pwwkew",
			},
			expected: 3,
		},
		{
			args: args{
				s: "",
			},
			expected: 0,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.s)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, lengthOfLongestSubstring(tc.args.s))
		})
	}
}
