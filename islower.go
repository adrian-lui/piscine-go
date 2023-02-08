package piscine

func IsLower(s string) bool {
	for _, match := range s {
		if match < 'a' || match > 'z' {
			return false
		}
	}
	return true
}
