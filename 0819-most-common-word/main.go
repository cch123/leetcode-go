package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(mostCommonWord("Bob. hIt, baLl", []string{"bob", "hit"}))
}

func mostCommonWord(paragraph string, banned []string) string {
	paraBytes := bytes.ToLower([]byte(paragraph))
	cur := []byte{}
	str2Cnt := map[string]int{}
	bannedMap := map[string]bool{}
	for i := 0; i < len(banned); i++ {
		bannedMap[banned[i]] = true
	}
	max, res := 0, ""

	for i := 0; i < len(paraBytes); i++ {
		if paraBytes[i] > 'z' || paraBytes[i] < 'a' {
			if len(cur) == 0 {
				continue
			}
			str2Cnt[string(cur)]++
			if str2Cnt[string(cur)] > max && bannedMap[string(cur)] == false {
				max = str2Cnt[string(cur)]
				res = string(cur)
			}
			cur = cur[:0]
		} else if i == len(paraBytes)-1 && len(cur) > 0 {
			cur = append(cur, paraBytes[i])
			str2Cnt[string(cur)]++
			if str2Cnt[string(cur)] > max && bannedMap[string(cur)] == false {
				max = str2Cnt[string(cur)]
				res = string(cur)
			}
		} else {
			cur = append(cur, paraBytes[i])
		}
	}
	fmt.Println(str2Cnt)
	return res
}
