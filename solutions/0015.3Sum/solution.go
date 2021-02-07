package pb0015

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i, x := range nums {
		if i > 0 && x == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			if nums[l]+nums[r] < -x {
				l++
			} else if nums[l]+nums[r] > -x {
				r--
			} else {
				res = append(res, []int{x, nums[l], nums[r]})
				l++
				r--
				for l < n && nums[l] == nums[l-1] {
					l++
				}
				for r >= 0 && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return res
}
