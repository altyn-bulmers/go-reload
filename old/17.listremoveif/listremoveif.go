package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	temp := l.Head
	prev := l.Head

	for temp != nil && temp.Data == data_ref {
		l.Head = temp.Next
		temp = l.Head
	}
	for temp != nil {
		for temp != nil && temp.Data != data_ref {
			prev = temp
			temp = temp.Next
		}

		if temp == nil {
			return
		}
		prev.Next = temp.Next
		temp = prev.Next	
	}
	
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}
