package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args[1:]
	length := 0
	var result int
	for range args {
		length++
	}

	if length != 3 {
		return
	}

	if args[1] != "+" && args[1] != "-" && args[1] != "/" && args[1] != "%" && args[1] != "*" {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	left := args[0]
	right := args[2]
	operator := args[1]

	if left[0] != '-' {
		number, _ := Atoi(left)
		if number < 0 {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
	}

	if right[0] != '-' {
		number, _ := Atoi(right)
		if number < 0 {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
	}

	// if left[0] == '-' && Atoi(left) > 0 {
	// 	z01.PrintRune('0')
	// 	z01.PrintRune('\n')
	// 	return
	// }

	if left[0] == '-' {
		number, _ := Atoi(left)
		if number > 0 {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
	}

	if right[0] == '-' {
		number, _ := Atoi(right)
		if number > 0 {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
	}

	// if right[0] == '-' && Atoi(right) > 0 {
	// 	z01.PrintRune('0')
	// 	z01.PrintRune('\n')
	// 	return
	// }

	left2, _ := Atoi(left)
	right2, _ := Atoi(right)
	_, left3 := Atoi(left)
	_, right3 := Atoi(right)

	switch operator {
	case "/":
		if right2 == 0 && right3 == true {
			output := "No division by 0"
			for _, x := range output {
				z01.PrintRune(x)
			}
			z01.PrintRune('\n')
			return
		} else if right2 == 0 && right3 == false {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		} else {
			result = left2 / right2
			if left2 < 0 && right2 < 0 && result < 0 {
				z01.PrintRune('0')
				z01.PrintRune('\n')
				return
			}
			//fmt.Println(result)

		}
	case "*":
		result = left2 * right2
		if left2 == 0 || right2 == 0 {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		} else if result/left2 != right2 || result/right2 != left2 {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
		//fmt.Println(result)
	case "+":
		result = left2 + right2
		if left2 > 0 && right2 > 0 && result < 0 || right3 == false || left3 == false || left2 < 0 && right2 < 0 && result > 0 {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
		//fmt.Println(result)
	case "-":
		result = left2 - right2
		if left2 > 0 && right2 < 0 && result < 0 || right3 == false || left3 == false || left2 < 0 && right2 > 0 && result > 0 {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
		//fmt.Println(result)
	case "%":
		if right2 == 0 && right3 == true {
			output := "No modulo by 0"
			for _, x := range output {
				z01.PrintRune(x)
			}
			z01.PrintRune('\n')
			return
		} else if right2 == 0 && right3 == false {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		} else {
			result = left2 % right2
			// if left2 < 0 && right2 < 0 && result < 0 {
			// 	z01.PrintRune('0')
			// 	z01.PrintRune('\n')
			// 	return
			// }
			//fmt.Println(result)

		}
	}

	finalres := Itoa(result)

	for _, final := range finalres {
		z01.PrintRune(final)
	}
	z01.PrintRune('\n')

}

//left2 > 0 && right2 > 0 && result < 0 || left2 < 0 && right2 < 0 && result < 0 || left2 < 0 && right2 > 0 && result > 0 || left2 > 0 && right2 < 0 && result > 0 ||

func Atoi(s string) (int, bool) {

	s0 := s
	var res int
	counts := 0
	for range s {
		counts++
	}

	if counts >= 1 && counts <= 20 {
		if s[0] == '-' || s[0] == '+' {
			s = s[1:]
		}
	} else {
		return 0, false
	}

	for _, v := range s {
		if v >= 48 && v <= 57 {
			res = int(v-48) + res*10
		} else {
			return 0, false
		}
	}
	if s0[0] == '-' {
		res = -res
	}

	return res, true
}

func Itoa(n int) string {

	var result string
	flag := false

	if n == -9223372036854775808 {
		return "-9223372036854775808"
	} else if n == 0 {
		return "0"
	}
	if n < 0 {
		n = -n
		flag = true

	}

	for n > 0 {
		d := (n % 10) + 48
		result = string(d) + result
		n = (n / 10)
	}

	if flag {
		result = "-" + result
	}
	return result

}
