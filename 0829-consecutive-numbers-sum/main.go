package main

func checkValid(N int, part int) (bool, int) {
	avg := N / part
	var first int
	if part%2 == 0 {
		first = avg - (part/2 - 1)
	} else {
		first = avg - part/2
	}

	if first <= 0 {
		return false, first
	}

	last := first + part - 1
	sum := (first + last) * part / 2
	if sum == N {
		return true, first
	}
	return false, first
}

func consecutiveNumbersSum(N int) int {
	sum := 0
	for i := 1; i <= N; i++ {
		valid, first := checkValid(N, i)
		if valid {
			sum++
		}
		if first <= 0 {
			break
		}
	}
	return sum
}
