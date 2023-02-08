package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 1 + i; j < len(a); j++ {
			if f(a[i], a[j]) == 1 {
				temp := a[j]
				a[j] = a[i]
				a[i] = temp
			}
		}
	}
}
