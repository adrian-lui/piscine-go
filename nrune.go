package piscine

func NRune(s string, n int) rune {
	if n > len(s) || n <= 0 {
		return 0
	} else {
		cut := []rune(s)
		return cut[n-1]
	}
}
