package piscine

func LastRune(s string) rune {
	cut := []rune(s)
	return cut[len(cut)-1]
}
