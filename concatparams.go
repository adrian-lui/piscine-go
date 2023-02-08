package piscine

func ConcatParams(args []string) string {
	var newStr string
	for i := 0; i < len(args); i++ {
		newStr += args[i]
		if i != len(args)-1 {
			newStr += "\n"
		}
	}
	return newStr
}
