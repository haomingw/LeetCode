package pb0007

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseInteger(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		args     args
		expected int
	}{
		{
			args: args{
				x: 123,
			},
			expected: 321,
		},
		{
			args: args{
				x: -123,
			},
			expected: -321,
		},
		{
			args: args{
				x: 120,
			},
			expected: 21,
		},
		{
			args: args{
				x: 1534236469,
			},
			expected: 0,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.x)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, reverse(tc.args.x))
		})
	}
}
