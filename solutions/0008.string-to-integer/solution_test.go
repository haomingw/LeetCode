package pb0008

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseInteger(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args     args
		expected int
	}{
		{
			args: args{
				s: "",
			},
			expected: 0,
		},
		{
			args: args{
				s: "42",
			},
			expected: 42,
		},
		{
			args: args{
				s: "   -42",
			},
			expected: -42,
		},
		{
			args: args{
				s: "4193 with words",
			},
			expected: 4193,
		},
		{
			args: args{
				s: "words and 987",
			},
			expected: 0,
		},
		{
			args: args{
				s: "-91283472332",
			},
			expected: -2147483648,
		},
		{
			args: args{
				s: "9223372036854775808",
			},
			expected: 2147483647,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.s)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, myAtoi(tc.args.s))
		})
	}
}
