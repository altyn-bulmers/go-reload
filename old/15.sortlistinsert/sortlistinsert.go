package piscine

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

