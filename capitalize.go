package piscine

func Capitalize(s string) string {
	var wordStart bool = true
	var newStr string
	var isAlpha bool
	for _, match := range s {
		if !(match >= 'a' && match <= 'z') && !(match >= 'A' && match <= 'Z') && !(match >= '0' && match <= '9') {
			isAlpha = false
		} else {
			isAlpha = true
		}
		if isAlpha {
			if match >= 'a' && match <= 'z' && wordStart {
				newStr += string(match - ' ')
			} else if match >= 'A' && match <= 'Z' && !wordStart {
				newStr += string(match + ' ')
			} else {
				newStr += string(match)
			}
			wordStart = false
		} else {
			newStr += string(match)
			wordStart = true
		}
	}
	return newStr
}
