package main

func main() {}
func moveZeroes(nums []int) {
	nonZeroPos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[nonZeroPos] = nums[nonZeroPos], nums[i]
			nonZeroPos++
		}
	}
}
