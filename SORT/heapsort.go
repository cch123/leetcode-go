package main

import "fmt"

func main() { arr := []int{9,8,7,6,5,5, 4, 3, 2, 1}; hs(arr); fmt.Println(arr) }

func hs(arr []int) {
	buildHeap(arr)

	for i := len(arr)-1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjustHeap(arr, 0, i-1)
	}
}

func buildHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		// 调整以 arr[i] 为堆顶的堆
		// left = 2*i+1, right = 2*i+2
		// biggest := arr[i]
		adjustHeap(arr, i, len(arr)-1)
	}
}

func adjustHeap(arr []int, minIdx int, maxIdx int) {
	i := minIdx
	for {
		lIdx, rIdx, idx, max := 2*i+1, 2*i+2, i, arr[i]
		if lIdx <= maxIdx && arr[lIdx] > max {
			idx, max = lIdx, arr[lIdx]
		}
		if rIdx <= maxIdx && arr[rIdx] > max {
			idx, max = rIdx, arr[rIdx]
		}
		if idx == i {break}
		arr[idx], arr[i] = arr[i], arr[idx]
		i = idx
	}
}

