package pb0003

func lengthOfLongestSubstring(s string) int {
	l, r, res := 0, 0, 0
	n := len(s)
	cnt := make([]int, 128)
	for l < n {
		if r < n && cnt[s[r]] == 0 {
			cnt[s[r]]++
			r++
		} else {
			cnt[s[l]]--
			l++
		}
		if r-l > res {
			res = r - l
		}
	}
	return res
}
