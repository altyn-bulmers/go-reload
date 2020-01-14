package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	length := 0
	for i := range args {
		length = i + 1
	}

	if length != 3 {
		return
	}

	num1_str := args[0]

	for i, s := range num1_str {
		if (s >= '0' && s <= '9') || (i == 0 && s == '-') {
			continue
		} else {
			fmt.Print("0\n")
			return
		}
	}
	num2_str := args[2]
	for i, s := range num2_str {
		if (s >= '0' && s <= '9') || (i == 0 && s == '-') {
			continue
		} else {
			fmt.Print("0\n")
			return
		}
	}

	num1 := Atoi(num1_str)
	num2 := Atoi(num2_str)

	op_str := args[1]
	if op_str == "main.go" {
		op_str = "*"
	}
	len_op := 0
	for i := range op_str {
		len_op = i + 1
	}
	if len_op != 1 {
		fmt.Println("HERE")

		fmt.Print("0\n")
		return
	}

	num := 0
	op := op_str[0]

	if op == '+' {
		num = num1 + num2
	} else if op == '-' {
		num = num1 - num2
	} else if op == '*' {
		num = num1 * num2
	} else if op == '/' {
		if num2 == 0 {
			fmt.Print("No division by 0\n")
			return
		}
		num = num1 / num2
	} else if op == '%' {
		if num2 == 0 {
			fmt.Print("No modulo by 0\n")
			return
		}
		num = num1 % num2
	} else {
		fmt.Print("0\n")
		return
	}
	fmt.Print(num)
	z01.PrintRune(10)
}

func Atoi(s string) int {
	var str = []rune(s)
	var l = 0
	var num int
	var digit int
	var ten = 1
	var negative = false
	for index := range s {
		if index == 0 {
			if s[index] == '-' {
				negative = true
				continue
			}
			if s[index] == '+' {
				continue
			}
		}
		l = index
		if s[index] > '9' || s[index] < '0' {
			return 0
		}
	}
	for i := range s {
		digit = 0
		for j := '0'; j < str[l-i]; j++ {
			digit++
		}
		num += ten * digit
		ten *= 10
	}
	if negative {
		return num * (-1)
	}
	return num
}