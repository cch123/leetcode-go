package main

func main() {}

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	diffPos := []int{}
	byteCnt := map[byte]int{}

	canSwapToTheSameFlag := false

	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			diffPos = append(diffPos, i)
		}
		byteCnt[A[i]]++
		if byteCnt[A[i]] == 2 {
			canSwapToTheSameFlag = true
		}
	}

	if len(diffPos) == 0 && canSwapToTheSameFlag {
		return true
	}

	if len(diffPos) != 2 {
		return false
	}

	x, y := diffPos[0], diffPos[1]
	if A[x] == B[y] && A[y] == B[x] {
		return true
	}

	return false
}
