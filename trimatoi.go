package piscine

func TrimAtoi(s string) int {
	neg := false
	var list []int
	var sum int = 0
	for _, match := range s {
		if match == '-' && len(list) == 0 {
			neg = true
		}
		if match >= '0' && match <= '9' {
			list = append(list, int(match))
		}
	}
	for index, num := range list {
		var multi int = 1
		for i := (len(list) - index - 1); i > 0; i-- {
			multi *= 10
		}
		sum += multi * (num - 48)
	}
	if neg == true {
		return sum * -1
	}
	return sum
}
