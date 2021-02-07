package pb0012

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		args     args
		expected string
	}{
		{
			args: args{
				num: 3,
			},
			expected: "III",
		},
		{
			args: args{
				num: 4,
			},
			expected: "IV",
		},
		{
			args: args{
				num: 9,
			},
			expected: "IX",
		},
		{
			args: args{
				num: 58,
			},
			expected: "LVIII",
		},
		{
			args: args{
				num: 1994,
			},
			expected: "MCMXCIV",
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.num)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, intToRoman(tc.args.num))
		})
	}
}
