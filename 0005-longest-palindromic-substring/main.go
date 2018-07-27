package main

func longestPalindrome(s string) string {
	length := len(s)
	var matrix = make([][]bool, length)
	for i := 0; i < length; i++ {
		matrix[i] = make([]bool, length)
	}

	var maxPalindromeLen, maxPalindrome = 0, ""

	// 初始化单字母和双字母是否回文的矩阵
	for i := 0; i < length; i++ {
		matrix[i][i] = true
		if maxPalindromeLen == 0 {
			maxPalindromeLen = 1
			maxPalindrome = string([]byte{s[i]})
		}
		if i > 0 && s[i] == s[i-1] {
			matrix[i-1][i] = true
			maxPalindromeLen = 2
			maxPalindrome = string([]byte{s[i-1], s[i]})
		}
	}

	// 在初始结果上，每次只可能增加两个字母
	// f(i,j) = f(i+1)(j-1) && s[i] == s[j]
	// 初始 i j = 0
	// 从矩阵的中轴线向两边扩展
	for substrLen := 3; substrLen <= length; substrLen++ {
		for i := 0; i < length; i++ {
			startIdx := i
			endIdx := i + substrLen - 1
			if endIdx >= length {
				break
			}

			if matrix[startIdx+1][endIdx-1] == true {
				if s[startIdx] == s[endIdx] {
					matrix[startIdx][endIdx] = true
					if substrLen > maxPalindromeLen {
						maxPalindromeLen = substrLen
						maxPalindrome = s[startIdx : endIdx+1]
					}
				}
			}
		}
	}

	return maxPalindrome
}
