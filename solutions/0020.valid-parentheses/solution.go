package pb0020

func isValid(s string) bool {
	a := []rune{}
	mp := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		if x, ok := mp[c]; ok {
			if len(a) == 0 || a[len(a)-1] != x {
				return false
			}
			a = a[:len(a)-1]
		} else {
			a = append(a, c)
		}
	}
	return len(a) == 0
}
