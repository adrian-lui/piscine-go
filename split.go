package piscine

func Split(s, sep string) []string {
	var newStrSlice []string
	var index int = 0
	for i := 0; i < len(s); i++ { // loop through s
		if i+len(sep) <= len(s) { // check sep and overflow
			if s[i:i+len(sep)] == sep { // check sep match
				newStrSlice = append(newStrSlice, s[index:i])
				index = i + len(sep)
			}
		} else {
			break
		}
	}
	newStrSlice = append(newStrSlice, s[index:])
	var count int = 0
	for i := len(newStrSlice) - 1; i >= 0; i-- { // remove empty
		if newStrSlice[i] == "" {
			count++
		} else {
			return newStrSlice[:len(newStrSlice)-count]
		}
	}
	return newStrSlice
}
