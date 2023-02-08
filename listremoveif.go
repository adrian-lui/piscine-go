package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	newList := &List{}
	iterator := l.Head
	for iterator != nil {
		if iterator.Data != data_ref {
			ListPushBack(newList, iterator.Data)
		}
		iterator = iterator.Next
	}
	*l = *newList
}
