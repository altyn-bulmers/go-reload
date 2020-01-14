package main

import (
//	"github.com/01-edu/z01.PrintRune"
	"fmt"
)

func main() {
	fmt.Println(RecursivePower(2,4))
}

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return RecursivePower(nb, power -1 ) * nb
}