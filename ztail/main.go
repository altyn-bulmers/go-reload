package main

import (
	"fmt"
	"os"
)

func main() {

	counter := 0

	var result []byte

	number, files, rev := numbers(os.Args[1:])

	for _, f := range files {
		if !rev {
			file, err := os.Open(f)
			if err != nil {
				fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", f)
				// if ind == ArrLen(files) - 1 {
				// 	fmt.Printf("\n")
				// }
				//os.Exit(0)
			} else {
				if ArrLen(files) > 1 {
					if counter == 0 {
						fmt.Printf("==> %s <==\n", f)
						counter++

					} else {
						fmt.Printf("\n==> %s <==\n", f)
					}
				}

				stat, err := file.Stat()
				if err != nil {
					fmt.Printf(err.Error())
					os.Exit(0)
				}
				size := stat.Size()
				sizeInt := int(size)
				if number > sizeInt {
					number = sizeInt
				}
				for i := 0; i < sizeInt; i++ {
					result = append(result, 0)
				}
				file.Read(result)
				fmt.Printf("%v", string(result[sizeInt-number:]))
			}

		} else {
			file, err := os.Open(f)
			if err != nil {
				fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", f)
				// if ind == ArrLen(files) - 1 {
				// 	fmt.Printf("\n")
				// }
				//os.Exit(0)
			} else {
				if ArrLen(files) > 1 {
					if counter == 0 {
						fmt.Printf("==> %s <==\n", f)
						counter++

					} else {
						fmt.Printf("\n==> %s <==\n", f)
					}
				}

				stat, err := file.Stat()
				if err != nil {
					fmt.Printf(err.Error())
					os.Exit(0)
				}
				size := stat.Size()
				sizeInt := int(size)
				if number > sizeInt {
					return
				}
				for i := 0; i < sizeInt; i++ {
					result = append(result, 0)
				}
				file.Read(result)
				fmt.Printf("%v", string(result[number-1:]))
			}

		}

	}
	//
	//fmt.Printf(rev)

}

func Atoi(s string) (int, bool) {

	// s0 := s
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
		return 1, false
	}

	for _, v := range s {
		if v >= 48 && v <= 57 {
			res = int(v-48) + res*10
		} else {
			return 0, false
		}
	}

	if res < 0 {
		return 1, false
	}
	// if s0[0] == '-' {
	// 	res = -res
	// }
	// if s[0] == '+' {
	// 	return res, true
	// }

	return res, true
}

func ArrLen(arr []string) int {
	count := 0
	for range arr {
		count++
	}
	return count
}

func strLen(arr string) int {
	count := 0
	for range arr {
		count++
	}
	return count
}

func numbers(args []string) (int, []string, bool) {

	cNum := 0
	//arg := os.Args[1:]
	ArgLength := 0
	var files []string
	//var result []byte
	flag := false

	for range args {
		ArgLength++
	}

	for index, ArgRange := range args {
		// var err bool
		// _, err = Atoi(ArgRange)
		// for _, i := range ArgRange {
		// 	if i == '+' {
		// 		flag = true
		// 	}
		// }
		if strLen(ArgRange) >= 2 && ArgRange[:2] == "-c" {
			if strLen(ArgRange) == 2 && index != ArgLength-1 {
				varNum, err := Atoi(args[index+1])

				for _, i := range ArgRange {
					if i == '+' {
						flag = true
					}
				}

				if varNum == 1 && err == false {
					fmt.Printf("tail: invalid number of bytes: ‘%s‘: Value too large for defined data type\n", args[index+1])
					os.Exit(1)
				}
				if varNum == 0 && err == false {
					fmt.Printf("tail: invalid number of bytes: ‘%s’\n", args[index+1])
					os.Exit(1)
				} else {
					cNum = varNum
					//flag = err
				}
			} else if strLen(ArgRange) > 2 {
				varNum, err := Atoi(ArgRange[2:])
				for _, i := range ArgRange {
					if i == '+' {
						flag = true
					}
				}

				if varNum == 1 && err == false {
					fmt.Printf("tail: invalid number of bytes: ‘%s‘: Value too large for defined data type\n", ArgRange[2:])
					os.Exit(1)
				}
				if varNum == 0 && err == false {
					fmt.Printf("tail: invalid number of bytes: ‘%s’\n", ArgRange[2:])
					os.Exit(1)
				} else {
					cNum = varNum
					for _, i := range ArgRange {
						if i == '+' {
							flag = true
						}
					}
					//flag = err
				}
			} else {
				fmt.Printf("tail: option requires an argument -- 'c'\n")
				fmt.Printf("Try 'tail --help' for more information.\n")
				os.Exit(0)
			}
			// if index >= ArgLength-1 {
			// 	fmt.Printf("tail: option requires an argument -- 'c'\n")
			// 	fmt.Printf("Try 'tail --help' for more information.\n")
			// 	os.Exit(0)
			// }

		} else {
			varNum, err := Atoi(ArgRange)
			for _, i := range ArgRange {
				if i == '+' {
					flag = true
				}
			}
			if ArgRange != "-c" && varNum == 0 && err == false {
				files = append(files, ArgRange)
			}
		}
		// return cNum, files, flag
	}
	return cNum, files, flag
}
