package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	new := &NodeI{Data: data_ref, Next: nil}

	if l == nil || l.Data > new.Data {
		new.Next = l
		return new
	} else {
		current := l
		for current.Next != nil && current.Next.Data < new.Data {
			current = current.Next
		}
		new.Next = current.Next
		current.Next = new
	}
	return l
}