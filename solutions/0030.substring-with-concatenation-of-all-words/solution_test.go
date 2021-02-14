package pb0030

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		args     args
		expected []int
	}{
		{
			args: args{
				s:     "barfoothefoobarman",
				words: []string{"foo", "bar"},
			},
			expected: []int{0, 9},
		},
		{
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "word"},
			},
			expected: []int{},
		},
		{
			args: args{
				s:     "barfoofoobarthefoobarman",
				words: []string{"bar", "foo", "the"},
			},
			expected: []int{6, 9, 12},
		},
		{
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "good"},
			},
			expected: []int{8},
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.s, tc.args.words)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, findSubstring(tc.args.s, tc.args.words))
		})
	}
}
