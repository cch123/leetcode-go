package main

func main() {}

type direction int

const (
	up direction = iota
	down
	unavail
)

func longestMountain(A []int) int {
	flag := unavail
	mountainLen := 0
	maxMountainLen := 0
	for i := 0; i < len(A)-1; i++ {
		switch flag {
		case unavail:
			if A[i+1] <= A[i] {
				// do nothing
			}
			if A[i+1] > A[i] {
				flag = up
				mountainLen++
			}
		case up:
			if A[i+1] < A[i] {
				flag = down
				mountainLen++
			}
			if A[i+1] > A[i] {
				mountainLen++
			}
			if A[i+1] == A[i] {
				mountainLen = 0
				flag = unavail
			}
		case down:
			if A[i+1] < A[i] {
				mountainLen++
			}
			if A[i+1] > A[i] {
				flag = up
				mountainLen++
				if maxMountainLen < mountainLen {
					maxMountainLen = mountainLen
				}
				mountainLen = 1
			}
			if A[i+1] == A[i] {
				flag = unavail
				mountainLen++
				if maxMountainLen < mountainLen {
					maxMountainLen = mountainLen
				}
				mountainLen = 0
			}
		}
	}
	if flag == down && A[len(A)-1] < A[len(A)-2] {
		mountainLen++
		if mountainLen > maxMountainLen {
			maxMountainLen = mountainLen
		}
	}
	return maxMountainLen
}
