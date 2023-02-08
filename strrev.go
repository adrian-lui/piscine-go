package piscine

func StrRev(s string) string {
	var x string
	for i := len(s) - 1; i >= 0; i-- {
		x = x + string(s[i])
	}
	return x
}
