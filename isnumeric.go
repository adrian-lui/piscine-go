package piscine

func IsNumeric(s string) bool {
	for _, match := range s {
		if match < '0' || match > '9' {
			return false
		}
	}
	return true
}
