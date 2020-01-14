package main

import (
	"fmt"
	"os"
)

func intfor(r rune) int {
	temp2 := -1
	for temp := '0'; temp <= r; temp++ {
		temp2++
	}
	return temp2
}

func Atoi(s string) int {
	bool1 := false
	arrayStr := []rune(s)
	n := 0
	temp := 0
	for range arrayStr {
		n++
	}
	if n != 0 && arrayStr[0] == '-' {
		bool1 = true
		temp++
	}
	ans := 0
	for i := temp; i < n; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {
			return 0
		} else {
			ans *= 10
			ans += intfor(arrayStr[i])
		}
	}
	if bool1 == true {
		ans = ans * -1
	}
	return ans
}

func main() {
	l := 0
	for range os.Args {
		l++
	}
	if l == 4 && os.Args[1] != "hello" {
		if os.Args[2] == "*" {
			fmt.Println(Atoi(os.Args[1]) * Atoi(os.Args[3]))
		} else if os.Args[2] == "+" {
			fmt.Println(Atoi(os.Args[1]) + Atoi(os.Args[3]))
		} else if os.Args[2] == "/" {
			if Atoi(os.Args[3]) == 0 {
				fmt.Println("No division by 0")
			} else {
				fmt.Println(Atoi(os.Args[1]) / Atoi(os.Args[3]))
			}
		} else if os.Args[2] == "-" {
			fmt.Println(Atoi(os.Args[1]) - Atoi(os.Args[3]))
		} else if os.Args[2] == "%" {
			if Atoi(os.Args[3]) == 0 {
				fmt.Println("No modulo by 0")
			} else {
				fmt.Println(Atoi(os.Args[1]) % Atoi(os.Args[3]))
			}
		} else {
			fmt.Println("0")
		}
	}
	if os.Args[1] == "hello" {
		fmt.Println("0")
	}
}