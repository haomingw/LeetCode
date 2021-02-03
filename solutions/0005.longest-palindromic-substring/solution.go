package pb0005

func longestPalindrome(s string) string {
	n, idx, sz := len(s), -1, -1
	var l, r int

	for i := 0; i < n; i++ {
		l, r = i, i
		for r < n-1 && s[r] == s[r+1] {
			r++
		}
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > sz {
			sz = r - l - 1
			idx = l + 1
		}
	}
	return s[idx : idx+sz]
}

func longestPalindrome2(s string) string {
	n, idx, sz := len(s), -1, -1
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = (s[i] == s[j]) && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > sz {
				idx, sz = i, j-i+1
			}
		}
	}
	return s[idx : idx+sz]
}
