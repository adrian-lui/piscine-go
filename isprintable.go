package piscine

func IsPrintable(s string) bool {
	for _, match := range s {
		if match < 32 || match > 126 {
			return false
		}
	}
	return true
}
