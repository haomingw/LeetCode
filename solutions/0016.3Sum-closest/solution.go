package pb0016

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n, diff := len(nums), 1<<30
	var res int
	for i, x := range nums {
		if i > 0 && nums[i-1] == x {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			t := x + nums[l] + nums[r]
			d := abs(target - t)
			if d < diff {
				diff, res = d, t
			}
			if t < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
