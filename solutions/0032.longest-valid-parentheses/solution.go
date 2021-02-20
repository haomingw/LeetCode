package pb0032

type pair struct {
	index int
	data  rune
}

func longestValidParentheses(s string) int {
	n, res := len(s), 0
	dp := make([]int, n+1)
	st := []pair{}
	for i, c := range s {
		if c == ')' && len(st) > 0 && st[len(st)-1].data == '(' {
			idx := st[len(st)-1].index
			dp[i+1] = dp[idx] + i - idx + 1
			st = st[:len(st)-1]
			if dp[i+1] > res {
				res = dp[i+1]
			}
		} else {
			st = append(st, pair{i, c})
		}
	}
	return res
}
