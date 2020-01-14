package main

import (
	"fmt"
)

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI{
	n := &NodeI{Data: data_ref, Next: nil}

	if l == nil || l.Data > n.Data {
		n.Next = l
		return n
	} else {
		curr := l
		for curr != nil && curr.Next.Data < n.Data {
			curr = curr.Next
		}
		n.Next = curr.Next
		curr.Next = n
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

func main() {

	var link *NodeI

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = SortListInsert(link, -2)
	link = SortListInsert(link, 2)
	PrintList(link)
}