package piscine

func BasicJoin(elems []string) string {
	var newWord string
	for _, word := range elems {
		newWord += word
	}
	return newWord
}
