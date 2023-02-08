package main

import (
	"os"
)

func main() {
	var num1 int
	var num2 int
	var result int
	var sign string

	if len(os.Args) != 4 {
		return
	}
	num1 = Atoi(os.Args[1])
	num2 = Atoi(os.Args[3])
	sign = os.Args[2]

	if num1 == 0 && os.Args[1] != "0" {
		return
	}
	if num2 == 0 && os.Args[3] != "0" {
		return
	}

	if num1 >= 9223372036854775807 || num1 <= -9223372036854775807 {
		return
	}

	if num2 >= 9223372036854775807 || num2 <= -9223372036854775807 {
		return
	}

	switch sign {
	case "+":
		result = num1 + num2
		if result < num1 || result < num2 {
			return
		}
	case "-":
		result = num1 - num2
		if result > num1 && result > num2 {
			return
		}
	case "*":
		result = num1 * num2
		if num1 == 9223372036854775807 || num2 == 9223372036854775807 {
			return
		}
	case "/":
		if num2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		} else {
			result = num1 / num2
		}
	case "%":
		if num2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		} else {
			result = num1 % num2
		}
	default:
		return
	}
	bytes := PrintNbr(result)
	os.Stdout.Write(bytes)
}

func PrintNbr(n int) []byte {
	var bytes []byte
	var runes []rune
	if n < 0 {
		bytes = append(bytes, '-')
	}
	if n == 0 {
		bytes = append(bytes, '0', '\n')
		return bytes
	} else {
		for {
			var remain int = n % 10
			if remain < 0 {
				remain = remain * -1
			}
			runes = append(runes, rune(remain))
			n = n / 10
			if n == 0 {
				break
			}
		}
	}
	for i := len(runes) - 1; i >= 0; i-- {
		bytes = append(bytes, byte(runes[i]+48))
	}
	bytes = append(bytes, '\n')
	return bytes
}

func Atoi(s string) int {
	var sum int = 0
	var signed bool = false
	var neg bool = false
	for index, x := range s {
		if x == '+' || x == '-' {
			if signed {
				return 0
			}
			if x == '-' {
				neg = true
			}
			signed = true
			continue
		}
		if x != '+' || x != '-' {
			if x > '9' || x < '0' {
				return 0
			}
		}
		digit := int(x) - 48
		power := (len(s) - index - 1)
		multiplier := 1
		for i := power; i > 0; i-- {
			multiplier *= 10
		}
		sum += digit * multiplier
		signed = true
	}
	if neg {
		sum *= -1
	}
	return sum
}
