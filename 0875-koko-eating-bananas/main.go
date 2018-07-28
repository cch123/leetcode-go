package main

func sumAndMax(nums []int) (int, int) {
	var sum, max = 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	return sum, max
}

func calcTimes(piles []int, K int) int {
	times := 0
	for _, pile := range piles {
		if pile%K == 0 {
			times += pile / K
		} else {
			times += pile/K + 1
		}
	}
	return times
}

func minEatingSpeed(piles []int, H int) int {
	sum, max := sumAndMax(piles)
	if H >= sum {
		return 1
	}

	lowerBound := sum / H
	upperBound := max
	minK := 99999999999

	for lowerBound <= upperBound {
		k := lowerBound + (upperBound-lowerBound)/2
		// > H 向右走
		// <=H 向左走
		times := calcTimes(piles, k)
		if times > H {
			lowerBound = k + 1
			continue
		}

		// calctimes <= H，记录当前的 K
		// 然后向左走
		if times <= H {
			upperBound = k - 1
			if k < minK {
				minK = k
			}
		}
	}

	return minK
}
