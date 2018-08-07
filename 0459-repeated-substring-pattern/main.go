package main

import "fmt"

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))
}

func match(s, substr string) bool {
	if len(s)%len(substr) != 0 || len(s) == len(substr) {
		return false
	}

	for i := 0; i < len(s); i += len(substr) {
		if s[i:i+len(substr)] != substr {
			return false
		}
	}
	return true
}

func repeatedSubstringPattern(s string) bool {
	for i := 1; i < len(s); i++ {
		substr := s[0:i]
		if match(s, substr) {
			return true
		}
	}
	return false
}
