package main

func plusOne(digits []int) []int {
	l := len(digits)
	pos := l - 1
	for pos >= 0 {
		digits[pos]++
		if digits[pos] <= 9 {
			return digits
		}

		digits[pos] = 0
		pos--
	}

	return append([]int{1}, digits...)
}
