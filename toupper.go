package piscine

func ToUpper(s string) string {
	var new_str string
	for _, match := range s {
		if match >= 'a' && match <= 'z' {
			new_str += string(match - ' ')
		} else {
			new_str += string(match)
		}
	}
	return new_str
}
