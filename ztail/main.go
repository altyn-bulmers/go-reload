package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	flagC := "-c"
	var c uint64
	var pm bool
	foundFlag := false
	mainpm := false
	err := ""
	arr1 := []string{}
	errMess := "tail: option requires an argument -- 'c'"
	errMess2 := "Try 'tail --help' for more information."
	errMess3 := "tail: option used in invalid context -- "
	errMess4 := "tail: invalid option -- ’"
	errMess5 := "tail: option '---' is ambiguous; possibilities: '---disable-inotify' '---presume-input-pipe'"
	errMess6 := "tail: unrecognized option '"
	if lenArrStr(args) == 1 {
		return
	}
	for i := 1; i < lenArrStr(args); {
		if args[i][0] == '-' {
			if args[i] == flagC {
				if i == lenArrStr(args)-1 {
					fmt.Printf("%s", errMess)
					z01.PrintRune(10)
					fmt.Printf("%s", errMess2)
					z01.PrintRune(10)
					return
				} else {
					foundFlag = true
					c, pm, err = atoiUint(args[i+1])
					if err != "" {
						printStrZ01(err)
						z01.PrintRune(10)
						return
					} else if pm {
						mainpm = true
					}
					i = i + 2
				}
			} else if args[i][1] == 'c' {
				foundFlag = true
				c, pm, err = atoiUint(args[i][2:])
				if err != "" {
					printStrZ01(err)
					z01.PrintRune(10)
					return
				} else if pm {
					mainpm = true
				}
				i++
			} else if args[i][1] != 'c' {
				if args[i][1] == '-' {
					if lenString(args[i]) == 2 {
						args = args[i+1:]
						for k := range args {
							arr1 = append(arr1, args[k])
						}
						break
					} else if args[i][2] == '-' {
						if lenString(args[i]) == 3 {
							fmt.Printf("%s", errMess5)
							z01.PrintRune(10)
							fmt.Printf("%s", errMess2)
							z01.PrintRune(10)
							return
						}
						fmt.Printf("%s%s’", errMess6, args[i])
						z01.PrintRune(10)
						fmt.Printf("%s", errMess2)
						z01.PrintRune(10)
						return
					}
				} else if args[i][1] >= '0' && args[i][1] <= '9' {
					fmt.Printf("%s", errMess3)
					z01.PrintRune(rune(args[i][1]))
					z01.PrintRune(10)
					return
				} else {
					fmt.Printf("%s%v’", errMess4, string(args[i][1]))
					z01.PrintRune(10)
					fmt.Printf("%s", errMess2)
					z01.PrintRune(10)
					return
				}

			}
		} else {
			arr1 = append(arr1, args[i])
			i++
		}
	}
	if foundFlag {
		for i := range arr1 {
			printArr(arr1[i], mainpm, c, lenArrStr(arr1), i)
		}
	} else {
		fmt.Printf("no flag '-c'")
		z01.PrintRune(10)
	}

}

func printArr(s string, pm bool, c uint64, len int, index int) {
	file, err := os.Open(s)
	if err != nil {
		fmt.Printf("tail: cannot open '%s' for reading: No such file or directory", s)
		z01.PrintRune(10)
		return
	}
	if len > 1 {
		if index != 0 {
			z01.PrintRune(10)
		}
		fmt.Printf("==> %s <==", s)
		z01.PrintRune(10)
	}
	f, _ := file.Stat()
	size := int64(f.Size())
	arr2 := []byte{}
	var i int64
	for i = 0; i < size; i++ {
		arr2 = append(arr2, 0)
	}
	file.Read(arr2)
	if pm {
		c--
		fmt.Printf("%s", string(arr2[c:]))
	} else {
		fmt.Printf("%v", string(arr2[lenArrByte(arr2)-c:]))
	}

}

func lenArrByte(arr []byte) uint64 {
	var l uint64
	l = 0
	for range arr {
		l++
	}
	return l
}
func lenArrStr(arr []string) int {
	l := 0
	for range arr {
		t := l
		l++
		if t > l {
			return l - 1
		}
	}
	return l
}
func lenString(s string) int {
	l := 0
	for range s {
		t := l
		l++
		if t > l {
			return l - 1
		}
	}
	return l
}
func atoiUint(s string) (uint64, bool, string) {
	error := "tail: invalid number of bytes: '" + s + "'"
	var nbr uint64
	bool := false
	len := 0
	for range s {
		len++
		if len > 2 {
			break
		}
	}

	for i := range s {
		if (s[i] == '+' && len == 1) || (s[i] == '-' && len == 1) {
			return 0, false, error
		} else if (s[i] == '+' && i != 0) || (s[i] == '-' && i != 0) {
			return 0, false, error
		} else if s[i] == '+' && i == 0 {
			bool = true
		} else if s[i] == '-' && i == 0 {
			continue
		} else if s[i] >= '0' && s[i] <= '9' {
			var temp uint64
			for j := '0'; j < rune(s[i]); j++ {
				temp++
			}
			temp2 := nbr
			nbr = nbr*10 + temp
			if nbr < temp2 {
				error = error + ": Value too large for defined data type"
				return 0, false, error
			}
		} else {
			return 0, false, error
		}
	}
	error = ""
	return nbr, bool, error
}
func printStrZ01(s string) {
	for i := range s {
		z01.PrintRune(rune(s[i]))
	}
}
