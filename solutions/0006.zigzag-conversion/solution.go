package pb0006

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	runes := []rune(s)
	rows := make([][]rune, numRows)
	down := false
	r := 0
	for _, c := range runes {
		rows[r] = append(rows[r], c)
		if r == 0 || r == numRows-1 {
			down = !down
		}
		if down {
			r++
		} else {
			r--
		}
	}
	res := []rune{}
	for i := range rows {
		res = append(res, rows[i]...)
	}
	return string(res)
}
