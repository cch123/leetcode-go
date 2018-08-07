package main

import "fmt"

func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
	fmt.Println(rotateString("abcde", "abced"))
}

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if (len(A) == 1 || len(A) == 0) && A != B {
		return false
	}

	for i := 0; i < len(A); i++ {
		A = A[1:] + A[:1]
		if A == B {
			return true
		}
	}

	return false
}
