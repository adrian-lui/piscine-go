package piscine

func SortIntegerTable(table []int) {
	sorted := false
	for !sorted {
		for index := range table {
			sorted = true
			if index == len(table)-1 {
				continue
			}
			if table[index] > table[index+1] {
				swap := table[index]
				table[index] = table[index+1]
				table[index+1] = swap
				sorted = false
			}
		}
		for index := range table {
			if index == len(table)-1 {
				break
			}
			if table[index] > table[index+1] {
				sorted = false
				break
			}
		}
	}
}
