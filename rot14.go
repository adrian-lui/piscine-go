package piscine

func Rot14(s string) string {
	var newStr []rune
	for _, each := range s {
		if !(each >= 'A' && each <= 'Z') && !(each >= 'a' && each <= 'z') {
			newStr = append(newStr, each)
			continue
		}
		for i := 0; i < 14; i++ {
			if each == 'Z' || each == 'z' {
				each -= 26
			}
			each++
		}
		newStr = append(newStr, each)
	}
	return string(newStr)
}
