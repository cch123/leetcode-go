package main

import "fmt"

func bs(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] < target {
			low = mid + 1
			continue
		}

		if arr[mid] > target {
			high = mid - 1
			continue
		}
		return mid
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(bs(arr, 3))
	fmt.Println(bs(arr, 1))
}
