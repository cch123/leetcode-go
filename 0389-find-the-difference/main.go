package main

// t 是先重排 s，再添加字母的结果
func findTheDifference(s string, t string) byte {
	var sMap = map[byte]int{}
	var tMap = map[byte]int{}
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
		tMap[t[i]]++
	}
	tMap[t[len(s)]]++
	for b, cnt := range tMap {
		if cnt-sMap[b] == 1 {
			return b
		}
	}

	return 0
}
