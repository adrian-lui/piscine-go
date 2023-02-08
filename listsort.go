package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	x := NodeI{Data: l.Data}
	x.Next = nil
	first := true

	for l != nil {
		if first {
			l = l.Next
			first = false
			continue
		}
		if x.Data > l.Data {
			temp := x
			x = *l
			x.Next = &temp
		} else {
			iterator := &x
			for iterator.Next != nil && l.Data > iterator.Next.Data {
				iterator = iterator.Next
			}
			y := NodeI{Data: l.Data}
			y.Next = iterator.Next
			iterator.Next = &y
		}
		l = l.Next
	}
	return &x
}
