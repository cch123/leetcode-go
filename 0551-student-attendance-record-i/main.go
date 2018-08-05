package main

func checkRecord(s string) bool {
	var currentChar byte
	var currentCharCnt int
	var totalAbsentCnt int
	for i := 0; i < len(s); i++ {
		if s[i] == currentChar {
			currentCharCnt++
		} else {
			currentChar = s[i]
			currentCharCnt = 1
		}

		if currentChar == 'L' && currentCharCnt > 2 {
			return false
		}

		if currentChar == 'A' {
			totalAbsentCnt++
			if totalAbsentCnt > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
}
