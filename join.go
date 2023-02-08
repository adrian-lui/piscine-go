package piscine

func Join(strs []string, sep string) string {
	var newStr string
	for index, word := range strs {
		newStr += word
		if index != (len(strs) - 1) {
			newStr += sep
		}
	}
	return newStr
}
