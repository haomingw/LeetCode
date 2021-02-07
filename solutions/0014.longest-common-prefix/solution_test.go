package pb0014

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		args     args
		expected string
	}{
		{
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			expected: "fl",
		},
		{
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			expected: "",
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.strs)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, longestCommonPrefix(tc.args.strs))
		})
	}
}
