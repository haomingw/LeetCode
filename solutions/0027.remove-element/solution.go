package pb0027

func removeElement(nums []int, val int) int {
	p := 0
	for _, x := range nums {
		if x != val {
			nums[p] = x
			p++
		}
	}
	return p
}
