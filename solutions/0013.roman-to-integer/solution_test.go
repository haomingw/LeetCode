package pb0013

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args     args
		expected int
	}{
		{
			args: args{
				s: "III",
			},
			expected: 3,
		},
		{
			args: args{
				s: "IV",
			},
			expected: 4,
		},
		{
			args: args{
				s: "IX",
			},
			expected: 9,
		},
		{
			args: args{
				s: "LVIII",
			},
			expected: 58,
		},
		{
			args: args{
				s: "MCMXCIV",
			},
			expected: 1994,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.s)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, romanToInt(tc.args.s))
		})
	}
}
