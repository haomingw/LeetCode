package pb0028

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	for i := 0; i <= n-m; i++ {
		ok := true
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
	return -1
}
