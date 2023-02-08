package piscine

func AlphaCount(s string) int {
	list := []byte(s)
	count := 0
	for _, match := range list {
		if match >= 'A' && match <= 'Z' {
			count++
		}
		if match >= 'a' && match <= 'z' {
			count++
		}
	}
	return count
}
