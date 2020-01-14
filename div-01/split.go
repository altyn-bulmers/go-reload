package piscine

func Split(str, charset string) []string {
	strAsRune := []rune(str)
	charsetAsRune := []rune(charset)
	lenStr := 0
	lenCharset := 0
	size := 0
	for i := range strAsRune {
		lenStr = i + 1
	}
	for i := range charsetAsRune {
		lenCharset = i + 1
	}
	if lenStr > 0 && lenCharset > 0 {
		size++
		for i := 0; i < lenStr; i++ {
			if str[i] == charset[0] {
				if i+lenCharset+1 < lenStr {
					k := 0
					for j := i; j < i+lenCharset; j++ {
						if strAsRune[j] == charsetAsRune[k] {
							k++
						}
					}
					if k == lenCharset {
						size++
					}
				}
			}
		}
	}
	answer := make([]string, size)
	word := ""
	index := 0
	if lenStr > 0 && lenCharset > 0 {
		for i := 0; i < lenStr; i++ {
			if str[i] == charset[0] {
				if i+lenCharset+1 < lenStr {
					k := 0
					for j := i; j < i+lenCharset; j++ {
						if strAsRune[j] == charsetAsRune[k] {
							k++
						}
					}
					if k == lenCharset {
						answer[index] = word
						i += lenCharset - 1
						index++
						word = ""
					} else {
						word += string(strAsRune[i])
					}
				} else {
					word += string(strAsRune[i])
				}
			} else {
				word += string(strAsRune[i])
			}
		}
		answer[size-1] = word
	}
	return answer
}
