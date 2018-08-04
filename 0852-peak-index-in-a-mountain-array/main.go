package main

import "fmt"

// binary
func peakIndexInMountainArray(A []int) int {
	low, high := 1, len(A)-2
	for low <= high {
		mid := low + (high-low)/2
		if A[mid] > A[mid-1] && A[mid] > A[mid+1] {
			return mid
		}

		if A[mid] > A[mid-1] && A[mid] < A[mid+1] {
			low = mid + 1
			continue
		}

		if A[mid] < A[mid-1] && A[mid] > A[mid+1] {
			high = mid - 1
			continue
		}
	}
	return 0
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{1, 2, 3, 1}))
}
