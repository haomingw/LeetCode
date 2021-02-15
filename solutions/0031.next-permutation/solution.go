package pb0031

func nextPermutation(nums []int) {
	n, mx, idx := len(nums), -1, -1
	for i := n - 1; i >= 0; i-- {
		if nums[i] > mx {
			mx = nums[i]
			idx = i
		} else {
			nums[i], nums[idx] = nums[idx], nums[i]
			break
		}
	}
	if idx == 0 {
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
