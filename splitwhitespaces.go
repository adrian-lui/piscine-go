package piscine

func SplitWhiteSpaces(s string) []string {
	var newStrSlice []string
	var index int = 0
	var isWord bool = false
	for i := 0; i < len(s); i++ {
		if (s[i] == ' ' || s[i] == '\n' || s[i] == '	') && isWord {
			word := s[index:i]
			newStrSlice = append(newStrSlice, word)
			index = i + 1
			isWord = false
		} else if (s[i] == ' ' || s[i] == '\n' || s[i] == '	') && !isWord {
			index++
		} else {
			isWord = true
		}
	}
	newStrSlice = append(newStrSlice, s[index:])
	return newStrSlice
}
