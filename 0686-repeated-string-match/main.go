package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("abc", "cabcabca"))
	fmt.Println(repeatedStringMatch("abcbc", "cabcbca"))
}
func repeatedStringMatch(A string, B string) int {
	startA := A
	maxLen := len(B) * 2
	if len(A) > len(B) {
		maxLen = 2 * len(A)
	}
	for i := 0; ; i++ {
		if strings.Contains(A, B) {
			return i + 1
		}
		if len(A) > maxLen {
			break
		}
		A += startA
	}
	return -1
}
