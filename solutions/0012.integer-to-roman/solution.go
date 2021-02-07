package pb0012

func intToRoman(num int) string {
	runes := []rune{}
	toRoman := map[int]rune{
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}
	base := 1000
	for num > 0 {
		x := num / base
		if x <= 3 {
			for i := 0; i < x; i++ {
				runes = append(runes, toRoman[base])
			}
		} else if x == 4 {
			runes = append(runes, toRoman[base])
			runes = append(runes, toRoman[base*5])
		} else if x <= 8 {
			runes = append(runes, toRoman[base*5])
			for i := 0; i < x-5; i++ {
				runes = append(runes, toRoman[base])
			}
		} else {
			runes = append(runes, toRoman[base])
			runes = append(runes, toRoman[base*10])
		}
		num -= x * base
		base /= 10
	}
	return string(runes)
}
