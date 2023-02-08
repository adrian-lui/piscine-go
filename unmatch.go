package piscine

func Unmatch(a []int) int {
	var foundMatch int = 0
	for _, match1 := range a {
		for _, match2 := range a {
			if match1 == match2 {
				foundMatch++
			}
		}
		if foundMatch%2 != 0 {
			return match1
		}
		foundMatch = 0
	}
	return -1
}
