package pb0022

func generateParenthesis(n int) []string {
	res := []string{}
	dfs(n, n, "", &res)
	return res
}

func dfs(numOpen int, numClose int, s string, res *[]string) {
	if numOpen == 0 && numClose == 0 {
		*res = append(*res, s)
		return
	}
	if numOpen > 0 {
		dfs(numOpen-1, numClose, s+"(", res)
	}
	if numClose > 0 && numOpen < numClose {
		dfs(numOpen, numClose-1, s+")", res)
	}
}
