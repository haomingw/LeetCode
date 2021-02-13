package pb0029

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		args     args
		expected int
	}{
		{
			args: args{
				dividend: 10,
				divisor:  3,
			},
			expected: 3,
		},
		{
			args: args{
				dividend: 7,
				divisor:  -3,
			},
			expected: -2,
		},
		{
			args: args{
				dividend: -2147483648,
				divisor:  -1,
			},
			expected: 2147483647,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.dividend, tc.args.divisor)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, divide(tc.args.dividend, tc.args.divisor))
		})
	}
}
