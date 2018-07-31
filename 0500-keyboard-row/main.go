package main

import (
	"fmt"
	"strings"
)

var byteMap = map[byte]int{
	'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1,
	'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
	'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2,
	'h': 2, 'j': 2, 'k': 2, 'l': 2,
	'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3,
	'n': 3, 'm': 3,
}

func findWords(words []string) []string {
	res := []string{}
outer:
	for i := 0; i < len(words); i++ {
		wordL := strings.ToLower(words[i])
		for j := 1; j < len(wordL); j++ {
			if byteMap[wordL[j]] != byteMap[wordL[j-1]] {
				continue outer
			}
		}
		res = append(res, words[i])
	}
	return res
}

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
