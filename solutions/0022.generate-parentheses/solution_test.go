package pb0022

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args     args
		expected []string
	}{
		{
			args: args{
				n: 3,
			},
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			args: args{
				n: 1,
			},
			expected: []string{"()"},
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.n)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, generateParenthesis(tc.args.n))
		})
	}
}
