package pb0009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	l, r := 1, 1
	for r <= x/10 {
		r *= 10
	}
	for l < r {
		if (x/l)%10 != (x/r)%10 {
			return false
		}
		l *= 10
		r /= 10
	}
	return true
}
