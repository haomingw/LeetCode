package pb0004

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		args     args
		expected float64
	}{
		{
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			expected: 2.,
		},
		{
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			expected: 2.5,
		},
		{
			args: args{
				nums1: []int{0, 0},
				nums2: []int{0, 0},
			},
			expected: 0.,
		},
		{
			args: args{
				nums1: []int{},
				nums2: []int{1},
			},
			expected: 1.,
		},
		{
			args: args{
				nums1: []int{2},
				nums2: []int{},
			},
			expected: 2.,
		},
		{
			args: args{
				nums1: []int{3, 4},
				nums2: []int{1, 2},
			},
			expected: 2.5,
		},
	}
	for _, tc := range tests {
		testname := fmt.Sprintf("%v %v", tc.args.nums1, tc.args.nums2)
		t.Run(testname, func(t *testing.T) {
			assert.Equal(t, tc.expected, findMedianSortedArrays(tc.args.nums1, tc.args.nums2))
		})
	}
}
