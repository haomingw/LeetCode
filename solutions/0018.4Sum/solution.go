package pb0018

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			l, r := j+1, n-1
			x := target - nums[i] - nums[j]
			for l < r {
				if nums[l]+nums[r] < x {
					l++
				} else if nums[l]+nums[r] > x {
					r--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < n && nums[l-1] == nums[l] {
						l++
					}
					for r >= 0 && nums[r+1] == nums[r] {
						r--
					}
				}
			}
		}
	}
	return res
}
