package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	fmt.Println(nums)
	if len(nums) == 1 {
		return nums[0]
	}

	if nums[0] < nums[1] {
		return nums[0]
	}

	if nums[len(nums)-1] > nums[len(nums)-2] {
		return nums[len(nums)-1]
	}

	// binary search
	low, high := 1, len(nums)-2
	for low <= high {
		mid := low + (high-low)/2
		println(low, high, mid)
		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}

		if mid%2 == 0 {
			if nums[mid] == nums[mid-1] {
				high = mid - 1
			}

			if nums[mid] == nums[mid+1] {
				low = mid + 1
			}
		} else {
			if nums[mid] == nums[mid-1] {
				low = mid + 1
			}

			if nums[mid] == nums[mid+1] {
				high = mid - 1
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))

}
