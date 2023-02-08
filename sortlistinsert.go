package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		empty := NodeI{Data: data_ref}
		l = &empty
		return l
	}
	// x := NodeI{Data: data_ref}
	if data_ref < l.Data {
		temp := NodeI{Data: l.Data}
		temp.Next = l.Next
		l.Data = data_ref
		l.Next = &temp
	} else {
		iterator := l
		for iterator.Next != nil && data_ref > iterator.Next.Data {
			iterator = iterator.Next
		}
		x := NodeI{Data: data_ref}
		x.Next = iterator.Next
		iterator.Next = &x
	}
	return l
}
