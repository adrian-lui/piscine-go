package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	iterator := l
	for count != pos {
		if iterator.Next == nil {
			return nil
		}
		iterator = iterator.Next
		count++
	}
	return iterator
}
