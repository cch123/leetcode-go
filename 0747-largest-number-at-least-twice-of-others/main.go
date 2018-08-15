package main

func main() {

}

type nu struct {
	num int
	idx int
}

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	if len(nums) == 0 {
		return -1
	}

	top1, top2 := nu{num: nums[0], idx: 0}, nu{num: nums[1], idx: 1}
	if top1.num < top2.num {
		top1, top2 = top2, top1
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > top1.num {
			top2 = top1
			top1 = nu{
				num: nums[i],
				idx: i,
			}
			continue
		}
		if nums[i] > top2.num {
			top2 = nu{
				num: nums[i],
				idx: i,
			}
		}
	}

	if top1.num >= top2.num*2 {
		return top1.idx
	}

	return -1
}
