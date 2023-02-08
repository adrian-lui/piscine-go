package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	iterator := n2
	for iterator != nil {
		SortListInsert(n1, iterator.Data)
		iterator = iterator.Next
	}
	return n1
}
