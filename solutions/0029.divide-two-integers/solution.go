package pb0029

func divide(dividend int, divisor int) int {
	if dividend == -(1<<31) && divisor == -1 {
		return 1<<31 - 1
	}
	a, b, res := abs(dividend), abs(divisor), 0
	for i := 31; i >= 0; i-- {
		if (a>>i)-b >= 0 {
			res += 1 << i
			a -= b << i
		}
	}
	if (dividend >= 0) != (divisor >= 0) {
		res *= -1
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
