package main

func main() {}

func checkPossibility(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}

	cnt := 0

	for i := 1; i < len(nums); i++ {
		// 波谷
		if nums[i] < nums[i-1] {
			cnt++
			if cnt > 1 {
				return false
			}

			// 如果前面的前面有值
			if i-2 >= 0 {
				if nums[i-2] <= nums[i] {
					// 说明改成 nums[i] 是安全的
					nums[i-1] = nums[i]
				} else {
					// 说明 nums[i-1] > nums[i-2] > nums[i]
					// 只能改 nums[i] = nums[i-1]
					nums[i] = nums[i-1]
				}
				continue
			}

			// 如果前面的前面没值
			if i-2 < 0 {
				nums[i-1] = nums[i]
				continue
			}
		}
	}

	return true
}
