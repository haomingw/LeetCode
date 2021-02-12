package pb0026

func removeDuplicates(nums []int) int {
	p := 0
	for i := range nums {
		if i == 0 || nums[i-1] != nums[i] {
			nums[p] = nums[i]
			p++
		}
	}
	return p
}
