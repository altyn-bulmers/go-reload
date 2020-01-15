package piscine

import (
	"fmt"
)

func PrintList(l *NodeI) {/////
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}