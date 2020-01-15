package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Div(a, b int) int {
	return a / b
}
func Mult(a, b int) int {
	return a * b
}
func Mod(a, b int) int {
	return a % b
}
func Plus(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}

type Calculator struct {
	symbol  string
	arrFunc func(int, int) int
}

func main() {
	arg := os.Args
	arguments := []string(arg[1:])
	lenArg := 0
	arrCalculator := []Calculator{
		{"/", Div},
		{"*", Mult},
		{"%", Mod},
		{"+", Plus},
		{"-", Minus},
	}
	for range arguments {
		lenArg++
	}
	result := 0
	errSym := -1
	if lenArg == 3 {
		if IsNumericD(arguments[0]) && IsNumericD(arguments[2]) {
			for i, fun := range arrCalculator {
				if arguments[1] == fun.symbol {
					errSym++
					if Atoi(arguments[2]) == 0 && i == 0 {
						Print("No division by 0\n")
						return
					} else if Atoi(arguments[2]) == 0 && i == 2 {
						Print("No modulo by 0\n")
						return
					} else {
						if Atoi(arguments[0]) == 0 && LenStr(arguments[0]) > 1 ||
							Atoi(arguments[2]) == 0 && LenStr(arguments[2]) > 1 {
							Print("0\n")
							return
						}
						result = fun.arrFunc(Atoi(arguments[0]), Atoi(arguments[2]))

						// check result
						var position int
						if fun.symbol == "+" {
							if Atoi(arguments[0]) > 0 && Atoi(arguments[2]) > 0 && result < 0 {
								Print("0\n")
								return
							}
							if Atoi(arguments[0]) < 0 && Atoi(arguments[2]) < 0 && result > 0 {
								Print("0\n")
								return
							}
						} else if fun.symbol == "-" {
							if Atoi(arguments[0]) > 0 && Atoi(arguments[2]) < 0 && result < 0 {
								Print("0\n")
								return
							}
							if Atoi(arguments[0]) < 0 && Atoi(arguments[2]) > 0 && result > 0 {
								Print("0\n")
								return
							}
						} else if fun.symbol == "*" {
							position = 0
							if Atoi(arguments[0]) != arrCalculator[position].arrFunc(result, Atoi(arguments[2])) {
								Print("0\n")
								return
							}
						} else if fun.symbol == "/" {
							position = 1
							if Atoi(arguments[0]) != arrCalculator[position].arrFunc(result, Atoi(arguments[2])) {
								Print("0\n")
								return
							}
						}
						break
					}
				}
			}
			if errSym == -1 {
				Print("0\n")
				return
			}
		}
		PrintInt(result)
		z01.PrintRune('\n')
	} else {
		Print("0\n")
		return
	}
}

func Print(s string) {
	for _, l := range []rune(s) {
		z01.PrintRune(l)
	}
}

func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	minus := ""
	if n < 0 {
		minus += "-"
		n *= -1
	}
	result := ""
	for n != 0 {
		num := n % 10
		counter := 0
		elem := '0'
		for i := '0'; i <= '9'; i++ {
			if counter == num {
				elem = i
				break
			}
			counter++
		}
		result = string(elem) + result
		n /= 10
	}
	if minus == "-" && result[0] == '0' {
		result = "0"
	} else if minus == "-" {
		result = minus + result
	}

	Print(result)
}

func IsNumericD(str string) bool {
	sAsRune := []rune(str)
	counter := 0
	strLen := LenStr(str)
	for _, letter := range sAsRune {
		if letter >= '0' && letter <= '9' {
			counter++
		}
	}
	if sAsRune[0] == '-' && strLen != 1 {
		counter++
	}
	if counter == strLen {
		return true
	}
	return false
}

func LenStr(s string) int {
	strLen := 0
	for range []rune(s) {
		strLen++
	}
	return strLen
}
func Atoi(s string) int {
	num := 0
	if !CheckStr(s) {
		return 0
	}
	strAsRune := []rune(s)
	minus := 1
	if strAsRune[0] == '+' || strAsRune[0] == '-' {
		if strAsRune[0] == '-' {
			minus = -1
		}
		strAsRune = strAsRune[1:]
	}
	for index, n := range strAsRune {
		if n != '0' {
			strAsRune = strAsRune[index:]
			for _, number := range strAsRune {
				counter := 0
				for i := '0'; i < number; i++ {
					counter++
				}
				//min -9,223,372,036,854,775,808; max 9,223,372,036,854,775,807
				if num > 922337203685477580 {
					return 0
				} else if num == 922337203685477580 {
					if minus == 1 {
						if counter > 8 {
							return 0
						}
					} else {
						if counter > 9 {
							return 0
						}
					}
					num = 10*num + counter
				} else {
					num = 10*num + counter
				}
			}
			break
		}
	}
	return num * minus
}

func CheckStr(s string) bool {
	if s == "" {
		return false
	}
	strAsRune := []rune(s)
	if strAsRune[0] == '+' || strAsRune[0] == '-' {
		strAsRune = strAsRune[1:]
	}
	if string(strAsRune) == "" {
		return false
	}
	numberLengh := 0
	for _, n := range strAsRune {
		if n < '0' || n > '9' {
			return false
		}
		numberLengh++
	}
	//min -9,223,372,036,854,775,808; max 9,223,372,036,854,775,807
	if numberLengh > 20 {
		return false
	}
	return true
}
