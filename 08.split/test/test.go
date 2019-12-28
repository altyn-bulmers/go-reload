package main

import (
	"fmt"
	student ".."
)

func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(student.Split(str, "HA"))
}
