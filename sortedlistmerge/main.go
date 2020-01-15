package main

import (
	"fmt"
)

type NodeI struct {
	Data int
	Next *NodeI
}

type List struct {
	Head *NodeI
	Tail *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1 = ListSort(n1)
	n2 = ListSort(n2)

	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Data <= n2.Data {
		n1.Next = SortedListMerge(n1.Next, n2)
		return n1
	} else {
		n2.Next = SortedListMerge(n1, n2.Next)
		return n2
	}
}

func ListSort(l *NodeI) *NodeI {

	if l != nil {
		NewH := l
		for NewH.Next != nil {
			NewH2 := l
			for NewH2.Next != nil {
				if NewH2.Next.Data < NewH2.Data {
					NewH2.Data, NewH2.Next.Data = NewH2.Next.Data, NewH2.Data
				}
				NewH2 = NewH2.Next
			}
			NewH = NewH.Next
		}
	}
	return l
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
