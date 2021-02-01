package pb0004

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	l, r, k := 0, len(nums1), (len(nums1)+len(nums2)+1)/2
	var x1, x2, y1, y2, mid int
	for l < r {
		mid = (l + r) / 2
		x1, x2 = get(nums1, mid-1), get(nums1, mid)
		y1, y2 = get(nums2, k-mid-1), get(nums2, k-mid)
		if x1 <= y2 && y1 <= x2 {
			l = mid
			break
		}
		if x1 > y2 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	x1, x2 = get(nums1, l-1), get(nums1, l)
	y1, y2 = get(nums2, k-l-1), get(nums2, k-l)
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(max(x1, y1))
	}
	return float64(max(x1, y1)+min(x2, y2)) / 2
}

func get(a []int, index int) int {
	if index < 0 {
		return minInt
	} else if index >= len(a) {
		return maxInt
	}
	return a[index]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
