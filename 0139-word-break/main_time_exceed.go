package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
}

func wordBreak(s string, wordDict []string) bool {
	if s == "" {
	}

	var res bool
	resArr := []bool{}
	for i := 0; i < len(wordDict); i++ {
		if len(wordDict[i]) <= len(s) && strings.HasPrefix(s, wordDict[i]) {
			curr := wordBreak(s[len(wordDict[i]):], wordDict)
			if curr == true {
				return true
			}
			resArr = append(resArr, curr)
		}
	}
	//fmt.Println(resArr, s, wordDict)
	for i := 0; i < len(resArr); i++ {
		res = res || resArr[i]
		if res == true {
			return res
		}
	}
	return res
}
