package pb0028

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		args     args
		expected int
	}{
		{
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			expected: 2,
		},
		{
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			expected: -1,
		},
		{
			args: args{
				haystack: "",
				needle:   "",
			},
			expected: 0,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.haystack, tc.args.needle)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, strStr(tc.args.haystack, tc.args.needle))
		})
	}
}
