package main

func pickOutIthElem(arr []int, i int) []int {
	if len(arr) == 0 {
		return arr
	}

	// copy arr to another slice
	res := make([]int, 0, len(arr)-1)
	for j := 0; j < len(arr); j++ {
		if j == i {
			continue
		}
		res = append(res, arr[j])
	}

	return res
}

func appendToEveryArr(arr [][]int, elem int) [][]int {
	var res [][]int
	for i := 0; i < len(arr); i++ {
		arr[i] = append(arr[i], elem)
		res = append(res, arr[i])
	}
	return res
}

func getPermutation(arr []int) [][]int {

	// 递归终止
	if len(arr) == 1 || len(arr) == 0 {
		return [][]int{arr}
	}

	var result = [][]int{}

	for i := 0; i < len(arr); i++ {
		// 去除 arr[i]
		arrWithOutIthElem := pickOutIthElem(arr, i)

		midResult := getPermutation(arrWithOutIthElem)

		// 给每一个子数组，都 append 上当前挑出的元素
		result = append(result, appendToEveryArr(midResult, arr[i])...)
	}
	return result
}

func permute(nums []int) [][]int {
	return getPermutation(nums)
}
