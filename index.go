package piscine

func Index(s string, toFind string) int {
	count := 0
	sRune := []rune(s)
	toFindRune := []rune(toFind)
	if len(toFindRune) == 0 {
		return 0
	}
	if len(sRune) == 0 {
		return -1
	}
	for i, match1 := range sRune {
		if match1 == rune(toFindRune[count]) {
			if count == len(toFindRune)-1 {
				return i - count
			}
			count++
		} else {
			count = 0
		}
	}
	return -1
}
