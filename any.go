package piscine

func Any(f func(string) bool, a []string) bool {
	for _, each := range a {
		if f(each) {
			return true
		}
	}
	return false
}
