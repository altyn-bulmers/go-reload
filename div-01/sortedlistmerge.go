package piscine

type NodeI struct {
	Data int
	Next *NodeI
}


func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	temp := n2
	for temp != nil {
		n1 = SortListInsert(n1, temp.Data)
		temp = temp.Next
	}
	return n1
}

// func SortListInsert(l *NodeI, data_ref int) *NodeI {
// 	if l == nil {
// 		return nil
// 	}
// 	if l.Data > data_ref { //if data_ref less than a begin node
// 		n := &NodeI{Data: data_ref, Next: l}
// 		return n
// 	}
// 	temp := l
// 	prevtemp := l
// 	for temp != nil {
// 		if temp.Data > data_ref {
// 			n := &NodeI{Data: data_ref, Next: temp}
// 			prevtemp.Next = n
// 			return l
// 		}
// 		prevtemp = temp
// 		temp = temp.Next
// 	}
// 	prevtemp.Next = &NodeI{Data: data_ref}
// 	return l
// }