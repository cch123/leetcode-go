package main

func main() {}

func shortestToChar(S string, C byte) []int {
	res := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			res[i] = 0
			continue
		}
		// left
		leftDis := len(S) + 1
		rightDis := len(S) + 1
		for j := i - 1; j >= 0; j-- {
			if S[j] == C {
				leftDis = i - j
				break
			}
		}
		// right
		for j := i + 1; j < len(S); j++ {
			if S[j] == C {
				rightDis = j - i
				break
			}
		}
		if leftDis < rightDis {
			res[i] = leftDis
		} else {
			res[i] = rightDis
		}
	}
	return res
}
