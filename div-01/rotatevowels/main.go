package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	sentance := ""
	arg := []string(arguments[1:])
	for index, word := range arg {
		if index == 0 {
			sentance = word
		} else {
			sentance += " " + word
		}
	}
	strAsRune := []rune(sentance)
	vowels := ""
	// 97-a
	// 101-e
	// 105-i
	// 111-o
	// 117-u
	for _, l := range strAsRune {
		if IsVowel(l) == true {
			vowels += string(l)
		}
	}
	if vowels != "" {
		vowels = StrRev(vowels)
	}
	vowelsAsRune := []rune(vowels)
	i := 0
	for index, l := range strAsRune {
		if IsVowel(l) == true {
			strAsRune[index] = vowelsAsRune[i]
			i++
		}
	}
	for _, letter := range strAsRune {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func IsVowel(l rune) bool {
	if l == 'a' || l == 'A' ||
		l == 'e' || l == 'E' ||
		l == 'i' || l == 'I' ||
		l == 'o' || l == 'O' ||
		l == 'u' || l == 'U' {
		return true
	}
	return false
}

func StrRev(s string) string {
	var count int
	for index := range s {
		count = index
	}
	j := 0
	revstr := []rune(s)
	for i := count; i >= 0; i-- {
		revstr[i] = rune(s[j])
		j++
	}
	finalstr := string(revstr)
	return finalstr
}
