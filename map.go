package piscine

func Map(f func(int) bool, a []int) []bool {
	var primeList []bool
	for _, each := range a {
		primeList = append(primeList, f(each))
	}
	return primeList
}
