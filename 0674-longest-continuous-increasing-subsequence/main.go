package main

func main() {}

func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	f := make([]int, len(nums))
	max := 1
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] <= nums[i-1] {
			f[i] = 1
			continue
		}

		if nums[i] > nums[i-1] {
			f[i] = f[i-1] + 1
		}

		if f[i] > max {
			max = f[i]
		}
	}
	return max
}
