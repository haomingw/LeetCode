package pb0017

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		args     args
		expected []string
	}{
		{
			args: args{
				digits: "23",
			},
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			args: args{
				digits: "",
			},
			expected: []string{},
		},
		{
			args: args{
				digits: "2",
			},
			expected: []string{"a", "b", "c"},
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.digits)
		t.Run(testname, func(t *testing.T) {
			actual := letterCombinations(tc.args.digits)
			sort.Strings(tc.expected)
			sort.Strings(actual)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
