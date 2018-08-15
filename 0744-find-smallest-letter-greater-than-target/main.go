package main

import "fmt"

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	// 记录比 target 小的最小
	// 记录比 target 大的最小
	var minLess, minBigger byte
	var minLessSet, minBiggerSet bool
	for i := 0; i < len(letters); i++ {
		if minLessSet {
			if letters[i] <= target && letters[i] <= minLess {
				minLess = letters[i]
			}
		} else {
			if letters[i] <= target {
				minLess = letters[i]
				minLessSet = true
			}
		}
		if minBiggerSet {
			if letters[i] > target && letters[i] < minBigger {
				minBigger = letters[i]
			}
		} else {
			if letters[i] > target {
				minBigger = letters[i]
				minBiggerSet = true
			}
		}
	}
	fmt.Println(minLessSet, minBiggerSet)
	fmt.Println(minLess, minBigger)
	if minBiggerSet {
		return minBigger
	}
	return minLess
}
