package pb0030

func findSubstring(s string, words []string) []int {
	mp, mpc := map[string]int{}, map[string]int{}
	res := []int{}
	n, m, k := len(s), len(words[0]), len(words)
	for i, w := range words {
		mp[w] = i
		mpc[w]++
	}
	a := make([]int, k)
	for i := 0; i < m; i++ {
		cnt := 0
		for j := range a {
			a[j] = 0
		}
		for j := i; j <= n-m; j += m {
			if val, ok := mp[s[j:j+m]]; ok {
				if a[val] < mpc[s[j:j+m]] {
					cnt++
				}
				a[val]++
			}
			if j-i >= k*m {
				if val, ok := mp[s[j-k*m:j-(k-1)*m]]; ok {
					a[val]--
					if a[val] < mpc[s[j-k*m:j-(k-1)*m]] {
						cnt--
					}
				}
			}
			if cnt == k {
				res = append(res, j-(k-1)*m)
			}
		}
	}
	return res
}
