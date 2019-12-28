package main

import  (
//"github.com/01-edu/z01"
"fmt"
)

func main() {
	string := "1234"
	fmt.Println(Atoi(string))
}


func Atoi(s string) int {
	
	runes := []rune(s)
	lenRune := 0
	result := 0
	for i := range runes {
		lenRune = i + 1
	}

	if lenRune == 0 {
		return 0
	}
	tens := 1
	for k := 0; k < lenRune -1; k++ {
		if runes[k] == '-' || runes[k] == '+' {
			continue
		}
		tens *= 10
	}

	for i := range runes {
		if (runes[0] == '-' || runes[0] == '+') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		numb := 0
		for r := '0'; r < runes[i]; r++ {
			numb++
		}
		result += tens * numb
		tens /= 10
	}

	if runes[0] == '-' {
		result *= -1
	}

	return result	

}