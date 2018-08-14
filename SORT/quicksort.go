package main

import "fmt"

func qs(arr []int) {
	if len(arr) == 0 || len(arr) == 1 {
		return
	}
	axis, pos := arr[0], len(arr)-1
	for i := pos; i >= 1; i-- {
		if arr[i] > axis {
			arr[i], arr[pos] = arr[pos], arr[i]
			pos--
		}
	}
	arr[0], arr[pos] = arr[pos], arr[0]
	qs(arr[:pos])
	qs(arr[pos+1:])
}

func main() {
	arr := []int{4, 3, 2, 1}
	qs(arr)
	fmt.Println(arr)
}
