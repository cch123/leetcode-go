package main

import "fmt"

func main() {
	arr := []int{4,3,2,1}
	ms(arr)
	fmt.Println(arr)
}

func ms(arr[]int) {
	if len(arr) == 0 || len(arr) == 1 { return }

	left,right := arr[:len(arr)/2], arr[len(arr)/2:]

	ms(left); ms(right)

	// merge
	x, y, idx, merged := 0,0,0, make([]int,len(arr))
	for x < len(left) && y < len(right) {
		if left[x] < right[y] {merged[idx] = left[x]; x++;idx++;continue;}
		if left[x] >= right[y] {merged[idx] = right[y]; y++;idx++;continue;}
	}
	for x < len(left) { merged[idx] = left[x]; idx++; x++ }
	for y < len(right) { merged[idx] = right[y]; idx++; y++ }
	copy(arr, merged)
}

