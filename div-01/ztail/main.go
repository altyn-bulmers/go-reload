package main

import (
	"fmt"
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	arguments := []string(arg[1:])
	size := 0
	lenArg := 0
	for i, input := range arguments {
		if input == "-c" {
			size = piscine.Atoi(arguments[i+1])
		}
		lenArg = i + 1
	}

	files := []string(arguments[2:])
	for _, filesN := range files {
		if lenArg > 3 {
			fmt.Printf("==> %v <==", filesN)
			fmt.Println()
		}

		file, err := os.Open(filesN)
		if err != nil {
			for _, l := range []rune(err.Error()) {
				z01.PrintRune(l)
			}
			z01.PrintRune('\n')
			os.Exit(0)
		}
		defer file.Close()

		fileStat, _ := os.Stat(filesN)
		sizeF := fileStat.Size()
		var content []byte
		for i := 0; i < int(sizeF); i++ {
			content = append(content, 0)
		}
		file.Read(content)

		lenF := 0
		for i := range content {
			lenF = i
		}
		if size <= lenF {
			content = content[lenF-size:]
			for i := range content {
				z01.PrintRune(rune(content[i]))
			}
			z01.PrintRune('\n')
		} else {
			for i := 0; i < lenF+1; i++ {
				z01.PrintRune(rune(content[i]))
			}
			z01.PrintRune('\n')
		}
	}
}
