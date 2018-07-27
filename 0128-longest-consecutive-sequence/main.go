package main

type bounds struct {
	lowerBound int
	upperBound int
}

func longestConsecutive(nums []int) int {
	var boundsMap = map[int]bounds{}
	var recordMap = map[int]bool{}
	var maxLen = 0

	for _, num := range nums {
		if recordMap[num] == true {
			// 已经处理过的应该跳过，要么可能影响后续的后并
			continue
		}

		left := num - 1
		right := num + 1
		newBounds := bounds{
			lowerBound: num,
			upperBound: num,
		}

		if _, ok := boundsMap[left]; ok {
			leftBounds := boundsMap[left]
			if leftBounds.upperBound >= num {
				continue
			}

			if leftBounds.upperBound == left {
				// 可以合并
				newBounds.lowerBound = leftBounds.lowerBound
				delete(boundsMap, left)
			}
		}

		if _, ok := boundsMap[right]; ok {
			rightBounds := boundsMap[right]
			if rightBounds.lowerBound <= num {
				continue
			}

			if rightBounds.lowerBound == right {
				// 可以合并
				newBounds.upperBound = rightBounds.upperBound
				delete(boundsMap, right)
			}
		}

		boundsMap[newBounds.lowerBound] = newBounds
		boundsMap[newBounds.upperBound] = newBounds
		currentLen := newBounds.upperBound - newBounds.lowerBound + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
		recordMap[num] = true
	}
	return maxLen
}
