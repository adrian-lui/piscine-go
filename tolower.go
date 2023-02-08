package piscine

func ToLower(s string) string {
	var new_str string
	for _, match := range s {
		if match >= 'A' && match <= 'Z' {
			new_str += string(match + ' ')
		} else {
			new_str += string(match)
		}
	}
	return new_str
}
