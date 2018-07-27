package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	lowered := strings.ToLower(s)

	filtered := ""

	for i := 0; i < len(lowered); i++ {
		if lowered[i] >= 'a' && lowered[i] <= 'z' {
			filtered += string([]byte{lowered[i]})
		}
		if lowered[i] >= '0' && lowered[i] <= '9' {
			filtered += string([]byte{lowered[i]})
		}
	}

	for i := 0; i < len(filtered)/2; i++ {
		if filtered[i] != filtered[len(filtered)-i-1] {
			return false
		}
	}
	return true
}
