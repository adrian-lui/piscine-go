package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	} else {
		newList := &List{}
		iterator := l.Head
		for iterator != nil {
			ListPushFront(newList, iterator.Data)
			iterator = iterator.Next
		}
		newList.Tail = l.Head
		newList.Tail.Next = nil
		*l = *newList
	}
}
