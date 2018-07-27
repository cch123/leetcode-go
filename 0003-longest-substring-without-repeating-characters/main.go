package main

func lengthfLongestSubstring(s string) int {
	var charMap = map[byte]bool{}
	var maxLen = 0
	var currentLen = 0

	for i := 0; i < len(s); i++ {
		// 如果 hashmap 中不存在
		// 那么说明这个元素可以无条件加入到无重复字母的字符串
		if charMap[s[i]] == true {
			// hashmap 中存在
			// 那要把一直到上一个重复字母(含)，之前的全部剔除掉
			loopIdx := i - currentLen
			for loopIdx < i {
				delete(charMap, s[loopIdx])
				currentLen--

				if s[loopIdx] == s[i] {
					break
				}
				loopIdx++
			}
		}

		charMap[s[i]] = true
		currentLen++
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}
