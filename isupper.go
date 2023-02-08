package piscine

func IsUpper(s string) bool {
	for _, match := range s {
		if match < 'A' || match > 'Z' {
			return false
		}
	}
	return true
}
