package main

import (
	"github.com/01-edu/z01"
//	"fmt"
	"os"
)

func printRune(array []rune) {
	for _ , r := range array{
		z01.PrintRune(r)
	}
}

func runeInSlice(a rune, list []rune) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func swap (vow []rune) {
	len := 0
	for i := range vow {
		len = i + 1
	}
	for i, j := 0, len -1; i < j; i, j = i +1, j -1 {
		temp := vow[i]
		vow[i] = vow[j]
		vow[j] = temp
	} 
}

func main(){
//	str := "HEllO World problem solved"
	argruments := os.Args[1:]
	length := 0
	for i := range argruments {
		length = i + 1
	}
	if length != 0 {
		str := ""
		first := false

		for _, arg := range argruments {
			if first {
				str += " "
			}
			first = true
			str += arg
		}
	

		vowels := []rune{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}
		length := 0
		vowLength := 0
		runes := []rune(str)

		for l := range runes {
			length = l + 1
		}

		for i:=0; i<length; i++{
			if runeInSlice(runes[i], vowels) {
				vowLength++
			}
		}

		vowArray := make([]rune, vowLength)
		pos := make([]int, vowLength)
		ind := 0
		for i, r := range runes {
			if runeInSlice(r, vowels) {
				pos[ind] = i
				vowArray[ind] = r
				ind++	
			}		
		}
		swap(vowArray)

		for i := range pos {
			runes[pos[i]] = vowArray[i]
		}

		printRune(runes)
	}
	z01.PrintRune('\n')
}
