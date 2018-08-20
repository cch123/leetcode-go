package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
}

func mark(s string, offset int, wordDict []string, marks []bool) {
	fmt.Println(s[offset:], offset, marks)
	if len(s) == 0 || offset >= len(s) {
		return
	}
	for i := 0; i < len(wordDict); i++ {
		if len(s[offset:]) >= len(wordDict[i]) && strings.HasPrefix(s[offset:], wordDict[i]) {
			if marks[offset+len(wordDict[i])-1] == true {
				// 这里是为了递归剪枝
				// 否则的话数量级会非常之大
				// 比如这里的例子：
				// 10 ^ len(s)
				continue
			}
			marks[offset+len(wordDict[i])-1] = true
			println(wordDict[i])
			mark(s, offset+len(wordDict[i]), wordDict, marks)
		}
	}
}

func wordBreak(s string, wordDict []string) bool {
	marks := make([]bool, len(s))
	mark(s, 0, wordDict, marks)
	return marks[len(s)-1]
}
