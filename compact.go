package piscine

func Compact(ptr *[]string) int {
	var newStr []string
	var count int = 0
	for _, nonZero := range *ptr {
		if nonZero != "" {
			newStr = append(newStr, nonZero)
			count++
		}
	}
	*ptr = newStr
	return count
}
