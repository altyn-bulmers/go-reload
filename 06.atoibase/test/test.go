package main

import (
	"fmt"
	student ".."
)

func main() {
	fmt.Println(student.AtoiBase("125", "0123456789"))
	fmt.Println(student.AtoiBase("1111101", "01"))
	fmt.Println(student.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(student.AtoiBase("uoi", "choumi"))
	fmt.Println(student.AtoiBase("bbbbbab", "-ab"))
}
