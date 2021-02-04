package pb0006

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZigZagConversion(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		args     args
		expected string
	}{
		{
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 1,
			},
			expected: "PAYPALISHIRING",
		},
		{
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			expected: "PAHNAPLSIIGYIR",
		},
		{
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			expected: "PINALSIGYAHRPI",
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.s, tc.args.numRows)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, convert(tc.args.s, tc.args.numRows))
		})
	}
}
