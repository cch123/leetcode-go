package main

func main() {
}

func isMonotonic(A []int) bool {
	if len(A) == 1 || len(A) == 0 {
		return true
	}
	monoFlag := 0

	for i := 1; i < len(A); i++ {
		if monoFlag == 0 {
			if A[i] > A[i-1] {
				monoFlag = 1 // 增
			}
			if A[i] < A[i-1] {
				monoFlag = 2
			}
			continue
		}
		if monoFlag == 2 && A[i] > A[i-1] {
			return false
		}
		if monoFlag == 1 && A[i] < A[i-1] {
			return false
		}
	}
	return true
}
