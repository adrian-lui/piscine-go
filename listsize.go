package piscine

func ListSize(l *List) int {
	var count int
	iterator := l.Head
	for iterator != nil {
		iterator = iterator.Next
		count++
	}
	return count
}
