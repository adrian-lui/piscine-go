package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int = 0
	for _, each := range tab {
		if f(each) {
			count++
		}
	}
	return count
}
