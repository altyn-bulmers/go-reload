package main

import (
	"fmt"
	student ".."
)

func main() {
	result := student.ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}