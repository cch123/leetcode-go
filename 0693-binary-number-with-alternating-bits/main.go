package main

func hasAlternatingBits(n int) bool {
	first := 1
	if n%2 == 0 {
		first = 2
	}

	for i := first; n > 0; i *= 4 {
		n -= i
	}

	if n == 0 {
		return true
	}

	return false

}
