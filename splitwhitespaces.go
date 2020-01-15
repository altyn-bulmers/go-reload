package piscine

func SplitWhiteSpaces(str string) []string {

	var n int
	var index int
	var count int
	var s string

	for j, rangeStr := range str {
		if rangeStr == 32 || rangeStr == 9 || rangeStr == 10 {
			s = str[n:j]
			n = j + 1
			if s != "" && s != "\n" {
				count++
			}
		}
	}

	n = 0
	arr := make([]string, count+1)
	for i, rangeStr := range str {
		if rangeStr == 32 || rangeStr == 9 || rangeStr == 10 {
			s = str[n:i]
			n = i + 1
			if s != "" && s != "\n" {
				arr[index] = s
				index++

			}

		} else if i == StrLen(str)-1 {
			if str[n-1:] != "" && str[n-1:] != "\n" {
				arr[index] = str[n:]
			}

		}
	}

	// for arr[len(arr)-1] == "" {
	// 	arr = arr[:len(arr)-1]
	// }
	lenarr := 0
	for range arr {
		lenarr++
	}

	if arr[lenarr-1] == "" {
		lenarr := 0
		for range arr {
			lenarr++
		}
		arr = arr[:lenarr-1]
	}
	return arr

}

func StrLen(str string) int {
	count := 0
	printword := []rune(str)
	for index := range printword {
		if index >= 0 {
			count++
		}
	}
	return count
}
