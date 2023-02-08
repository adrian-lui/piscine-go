package piscine

func ForEach(f func(int), a []int) {
	for _, each := range a {
		f(each)
	}
}
