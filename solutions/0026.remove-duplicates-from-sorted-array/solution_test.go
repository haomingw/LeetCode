package pb0026

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args     args
		expected []int
	}{
		{
			args: args{
				nums: []int{1, 1, 2},
			},
			expected: []int{1, 2},
		},
		{
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			expected: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v", tc.args.nums)
		t.Run(testname, func(t *testing.T) {
			end := removeDuplicates(tc.args.nums)
			assert.Equal(t, tc.expected, tc.args.nums[:end])
		})
	}
}
