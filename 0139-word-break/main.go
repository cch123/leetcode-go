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
	//fmt.Println(s, offset, marks)
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(wordDict); i++ {
		if len(s) >= len(wordDict[i]) && strings.HasPrefix(s, wordDict[i]) {
			if marks[offset+len(wordDict[i])-1] == true {
				continue
			}
			marks[offset+len(wordDict[i])-1] = true
			//println(wordDict[i])
			mark(s[len(wordDict[i]):], offset+len(wordDict[i]), wordDict, marks)
		}
	}
}

func wordBreak(s string, wordDict []string) bool {
	marks := make([]bool, len(s))
	mark(s, 0, wordDict, marks)
	//fmt.Println(marks)
	return marks[len(s)-1]
}
