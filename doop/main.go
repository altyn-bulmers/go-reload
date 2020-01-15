package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lens := 0
	for range os.Args {
		lens++
	}
	if lens != 4 {
		return
	}
	var a, b int64
	var answer string
	str1 := os.Args[1]
	str2 := os.Args[3]

	if isNumeric(str1) == false || isNumeric(str2) == false {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	a = Atoi(str1)
	b = Atoi(str2)
	operator := os.Args[2]

	if str1[0] != '-' && a < 0 || str1[0] == '-' && a > 0 || str2[0] != '-' && b < 0 || str2[0] == '-' && b > 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	if operator == "+" {

		if a > 0 && b > 0 && a+b < 0 || a < 0 && b < 0 && a+b > 0 {
			answer = Itoa1(0)
		}
		answer = Itoa1(a + b)
	} else if operator == "-" {

		if a < 0 && b > 0 && a-b > 0 || a > 0 && b < 0 && a-b < 0 {
			answer = Itoa1(0)
		}
		answer = Itoa1(a - b)

	} else if operator == "*" {
		answer = mult(a, b)

	} else if operator == "/" {

		if a < 0 && b < 0 && a/b < 0 {
			answer = Itoa1(0)
		}
		if b == 0 {
			answer = "No division by 0"
		} else {
			answer = Itoa1(a / b)
		}
	} else if operator == "%" {
		if b == 0 {
			answer = "No modulo by 0"
		} else {
			answer = Itoa1(a % b)
		}

	} else {

		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	for _, v := range answer {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func Itoa1(n int64) string {
	var s string
	sign := true

	if n == -9223372036854775808 {
		s = "0"
		return s
	}
	if n < 0 {
		n = n * -1
		sign = false
	}

	for {
		mod := n % 10
		s = string(mod+'0') + s
		n = n / 10
		if n == 0 {
			break
		}
	}

	if sign == false {
		s = "-" + s
	}
	return s
}

func mult(a, b int64) string {
	res := a * b

	if res/a != b {
		return Itoa1(0)
	}
	return Itoa1(a * b)
}

func isNumeric(str string) bool {

	s := []rune(str)

	for i := 0; i < StrLen(str); i++ {
		if !IsNbr(s[i]) {
			return false
		}
	}
	return true

}

func IsNbr(s rune) bool {
	if s >= '0' && s <= '9' || s == '-' {
		return true
	}
	return false

}

func StrLen(str string) int {
	var result int
	ByteStr := []rune(str)

	for index := range ByteStr {
		result = index + 1
	}

	return result
}

func Atoi(s string) int64 {
	var result int64
	var letter rune
	var sign int64 = 1

	for index := range s {
		letter = rune(s[index])
		if letter == '-' && index == 0 {
			sign *= -1
		} else if letter == '+' && index == 0 {
			continue
		} else if letter < '0' || letter > '9' {
			return 0
		} else {
			var n int64 = 0
			for i := '0'; i < letter; i++ {
				n++
			}
			result = result*10 + n
		}
	}

	return result * sign
}
