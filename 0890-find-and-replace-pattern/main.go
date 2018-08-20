package main

import "fmt"

func main() {
	fmt.Println(findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
}

func findAndReplacePattern(words []string, pattern string) []string {
	res := []string{}
outer:
	for i := 0; i < len(words); i++ {
		if len(words[i]) != len(pattern) {
			continue
		}
		m := map[byte]byte{}
		m2 := map[byte]byte{}
		word := words[i]
		for j := 0; j < len(word); j++ {
			if m[word[j]] == 0 && m2[pattern[j]] == 0 {
				m[word[j]] = pattern[j]
				m2[pattern[j]] = word[j]
			}
			if m[word[j]] == 0 && m2[pattern[j]] != 0 {
				continue outer
			}
			if m[word[j]] != 0 && m2[pattern[j]] == 0 {
				continue outer
			}
			if m[word[j]] != 0 && m2[pattern[j]] != 0 {
				if m[word[j]] != pattern[j] || m2[pattern[j]] != word[j] {
					continue outer
				}
			}
		}
		res = append(res, word)
	}
	return res
}
