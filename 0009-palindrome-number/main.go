package main

func isPalindrome(x int) bool {

	if x == 0 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}

	digits := []int{}
	y := x
	for y > 0 {
		digits = append(digits, y%10)
		y = y / 10
	}

	newSum := 0
	for i := 0; i < len(digits); i++ {
		newSum = newSum*10 + digits[i]
	}

	if newSum == x {
		return true
	}

	return false

}
