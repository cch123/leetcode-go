package main

import "fmt"

func decodeAtIndex(S string, K int) string {
	var base = 0
	var exp = 1
	var lastChar byte
	for i := 0; i < len(S); i++ {
		if S[i] >= 'a' && S[i] <= 'z' {
			base = base*exp + 1
			exp = 1
			lastChar = S[i]
		}
		if S[i] >= '0' && S[i] <= '9' {
			exp = (int(S[i]) - int('0')) * exp
		}
		//println(base, exp, S, K)
		if base*exp == K || (base*exp > K && K%base == 0) {
			return string([]byte{lastChar})
		}

		if base*exp > K {
			return decodeAtIndex(S, K%(base))
		}
	}
	return ""
}

func main() {
	fmt.Println(decodeAtIndex("leet2code3", 10))

	fmt.Println(decodeAtIndex("ha22", 5))
	fmt.Println(decodeAtIndex("a2345678999999999999999", 1))
	fmt.Println(decodeAtIndex("a23", 6))
	fmt.Println(decodeAtIndex("vk6u5xhq9v", 554))
}
