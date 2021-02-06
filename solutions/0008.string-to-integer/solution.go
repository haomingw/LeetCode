package pb0008

func myAtoi(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	runes := []rune(s)
	i, res := 0, 0
	negative := false
	for ; i < n && runes[i] == rune(' '); i++ {
	}
	if i < n && (runes[i] == rune('+') || runes[i] == ('-')) {
		negative = runes[i] == rune('-')
		i++
	}
	maxInt := 1<<31 - 1
	minInt := -maxInt - 1
	for ; i < n; i++ {
		if val, ok := toInt(runes[i]); ok {
			res = res*10 + val
			if res > maxInt {
				break
			}
		} else {
			break
		}
	}
	if negative {
		res *= -1
	}
	if res > maxInt {
		return maxInt
	}
	if res < minInt {
		return minInt
	}
	return res
}

func toInt(c rune) (int, bool) {
	res := int(c - rune('0'))
	if res >= 0 && res <= 9 {
		return res, true
	}
	return 0, false
}
