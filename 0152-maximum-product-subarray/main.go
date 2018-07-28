package main

func maxAndMin(sl ...int) (int, int) {
	var max = sl[0]
	var min = sl[0]
	for _, s := range sl {
		if s > max {
			max = s
		}
		if s < min {
			min = s
		}
	}
	return max, min
}

// 最大
// 要么就是当前元素乘前面的最小值(当都为负时得到)
// 要么就是当前元素乘前面的最大值(当都为正时得到)
// 要么就是当前元素，当之前的值 * 当前的值还不如只保留当前值更大时
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var res = nums[0]
	var max = nums[0]
	var min = nums[0]

	for i := 1; i < len(nums); i++ {
		max, min = maxAndMin(nums[i], nums[i]*max, nums[i]*min)
		if max > res {
			res = max
		}
	}

	return res
}
