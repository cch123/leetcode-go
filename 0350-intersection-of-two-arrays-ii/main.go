package main

func main() {
}

func intersect(nums1 []int, nums2 []int) []int {
	num1Map := map[int]int{}
	num2Map := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		num := nums1[i]
		num1Map[num]++
	}
	for i := 0; i < len(nums2); i++ {
		num := nums2[i]
		if _, ok := num1Map[num]; ok {
			num2Map[num]++
			if num2Map[num] > num1Map[num] {
				num2Map[num] = num1Map[num]
			}
		}
	}
	res := []int{}
	for num, cnt := range num2Map {
		for i := 0; i < cnt; i++ {
			res = append(res, num)
		}
	}
	return res
}
