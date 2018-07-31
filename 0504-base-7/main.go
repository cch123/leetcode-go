package main

import "fmt"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	negFlag := false
	if num < 0 {
		num = -num
		negFlag = true
	}

	var res = ""

	for num != 0 {
		mod := num % 7
		res = fmt.Sprint(mod) + res
		num /= 7
	}

	if negFlag {
		return "-" + res
	}

	return res
}

func main() {

	println(convertToBase7(100))
}
