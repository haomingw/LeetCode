package pb0013

func romanToInt(s string) int {
	res := 0
	toInt := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := range s {
		if i < len(s)-1 && toInt[s[i]] < toInt[s[i+1]] {
			res -= toInt[s[i]]
		} else {
			res += toInt[s[i]]
		}
	}
	return res
}
