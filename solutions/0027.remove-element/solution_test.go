package pb0027

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		args     args
		expected []int
	}{
		{
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			expected: []int{2, 2},
		},
		{
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4},
				val:  2,
			},
			expected: []int{0, 1, 3, 0, 4},
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.nums, tc.args.val)
		t.Run(testname, func(t *testing.T) {
			end := removeElement(tc.args.nums, tc.args.val)
			assert.Equal(t, tc.expected, tc.args.nums[:end])
		})
	}
}
