package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	iterator := l.Head
	lastData := l.Head.Data
	for iterator != nil {
		lastData = iterator.Data
		iterator = iterator.Next
	}
	return lastData
}
