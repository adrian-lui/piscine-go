package piscine

func IsAlpha(s string) bool {
	for _, match := range s {
		if !(match >= 'a' && match <= 'z') && !(match >= 'A' && match <= 'Z') && !(match >= '0' && match <= '9') {
			return false
		}
	}
	return true
}
