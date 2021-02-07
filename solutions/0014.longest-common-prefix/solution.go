package pb0014

func longestCommonPrefix(strs []string) string {
	var res []byte
	m := 1 << 31
	for _, s := range strs {
		if len(s) < m {
			m = len(s)
		}
	}
	for j := 0; j < m; j++ {
		ok := true
		for i := 1; i < len(strs); i++ {
			if strs[i-1][j] != strs[i][j] {
				ok = false
				break
			}
		}
		if !ok {
			break
		}
		res = append(res, strs[0][j])
	}
	return string(res)
}
