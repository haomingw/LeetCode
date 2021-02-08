package pb0017

var toLetters = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	var res = []string{}
	if len(digits) == 0 {
		return res
	}
	dfs(0, &digits, "", &res)
	return res
}

func dfs(index int, digits *string, s string, res *[]string) {
	if index == len(*digits) {
		*res = append(*res, s)
		return
	}
	digit := (*digits)[index]
	letters := toLetters[int(digit-'0')]
	for _, c := range letters {
		dfs(index+1, digits, s+string(c), res)
	}
}
