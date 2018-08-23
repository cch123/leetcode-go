package main

import "fmt"

func main() { arr := []int{4, 3, 2, 1}; hs(arr); fmt.Println(arr) }

func hs(arr []int) {
	if len(arr) == 0 || len(arr) == 1 {
		return
	}
	buildHeap(arr)

	for i := len(arr); i > 0; i-- {
		arr[0], arr[i-1] = arr[i-1], arr[0]
		adjustHeap(arr, 0, i-2)
	}
}

func buildHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		// 调整以 arr[i] 为堆顶的堆
		// left = 2*i+1, right = 2*i+2
		// smallest := arr[i]
		adjustHeap(arr, i, len(arr)-1)
	}
}

func adjustHeap(arr []int, minIdx int, maxIdx int) {
	j := minIdx
	for {
		smallest, idx := arr[j], j
		if 2*j+1 <= maxIdx && arr[2*j+1] > smallest {
			idx, smallest = 2*j+1, arr[2*j+1]
		}
		if 2*j+2 <= maxIdx && arr[2*j+2] > smallest {
			idx, smallest = 2*j+2, arr[2*j+1]
		}
		if idx == j {
			break
		}
		arr[j], arr[idx] = arr[idx], arr[j]
		j = idx
	}
}

