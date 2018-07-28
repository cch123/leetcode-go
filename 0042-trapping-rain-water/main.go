package main

func copyAndSum(height []int) ([]int, int) {
	dstArr := make([]int, len(height))
	sum := 0
	for i := 0; i < len(height); i++ {
		dstArr[i] = height[i]
		sum += height[i]
	}
	return dstArr, sum
}

func sum(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		res += height[i]
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func trap(height []int) int {
	dstArr, originalSum := copyAndSum(height)

	for i := 0; i < len(height); i++ {
		// println(i)
		if height[i] != 0 {
			// find the first that >= height[i]
			// if there is not such one
			// return the right max
			rightMax, idx, found := 0, 0, false
			for j := i + 1; j < len(height); j++ {
				if height[j] > rightMax {
					rightMax = height[j]
					idx = j
					found = true
				}
				if height[j] >= height[i] {
					rightMax = height[j]
					idx = j
					found = true
					break
				}
			}

			if !found {
				continue
			}
			// println(rightMax, height[i])
			minimum := min(rightMax, height[i])
			for k := i + 1; k < idx; k++ {
				dstArr[k] = minimum
			}
			i = idx - 1
		}
	}
	// fmt.Println(dstArr)

	newSum := sum(dstArr)
	return newSum - originalSum
}
