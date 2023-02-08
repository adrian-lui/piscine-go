package piscine

func ListMerge(l1 *List, l2 *List) {
	iterator := l2.Head

	for iterator != nil {
		ListPushBack(l1, iterator.Data)
		iterator = iterator.Next
	}
}
