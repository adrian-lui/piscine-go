package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a)-1; i++ {
		for j := 1 + i; j < len(a); j++ {
			if a[i] > a[j] {
				temp := a[j]
				a[j] = a[i]
				a[i] = temp
			}
		}
	}
}
