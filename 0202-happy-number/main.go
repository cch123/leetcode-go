package main

func main() {
}

func next(n int) int {
	sum := 0
	for n > 0 {
		mod := n % 10
		sum += mod * mod
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	seenBefore := map[int]bool{}
	for {
		if n == 1 {
			return true
		}

		if seenBefore[n] == true {
			return false
		}

		seenBefore[n] = true
		n = next(n)
	}
}
