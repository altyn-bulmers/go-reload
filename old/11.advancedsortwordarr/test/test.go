package main

import (
	"fmt"
	student ".."
)

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	student.AdvancedSortWordArr(result, student.Compare)

	fmt.Println(result)
}