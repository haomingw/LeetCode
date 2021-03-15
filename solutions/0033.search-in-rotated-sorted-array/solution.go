package pb0033

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[l] < nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid+1] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
