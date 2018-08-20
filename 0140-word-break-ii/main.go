package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

var res = [][]string{}

func dfs(s string, offset int, wordDict []string, current []string) {
	if offset == len(s) {
		res = append(res, current)
		return
	}

	for i := 0; i < len(wordDict); i++ {
		if len(s[offset:]) >= len(wordDict[i]) && strings.HasPrefix(s[offset:], wordDict[i]) {
			path := make([]string, len(current)+1)
			copy(path, current)
			path[len(path)-1] = wordDict[i]
			dfs(s, offset+len(wordDict[i]), wordDict, path)
		}
	}
}

// dfs
func wordBreak(s string, wordDict []string) []string {
	if !wordBreak1(s, wordDict) {
		return []string{}
	}
	res = res[:0]
	dfs(s, 0, wordDict, []string{})
	final := []string{}
	fmt.Println(res)
	for i := 0; i < len(res); i++ {
		final = append(final, strings.Join(res[i], " "))
	}
	return final
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

func wordBreak1(s string, wordDict []string) bool {
	marks := make([]bool, len(s))
	mark(s, 0, wordDict, marks)
	return marks[len(s)-1]
}
