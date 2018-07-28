package main

import "fmt"

var res = []string{"1", "1"} //"11", "21", "1211", "111221"}

func countInt(n int) string {
	countStr := ""
	nStr := res[n-1]
	for i := 0; i < len(nStr); i++ {
		cnt := 1
		currentChar := nStr[i]
		for i+1 < len(nStr) && nStr[i+1] == nStr[i] {
			cnt++
			i++
		}
		countStr = countStr + fmt.Sprint(cnt) + string([]byte{currentChar})
	}

	return countStr
}

func countAndSay(n int) string {
	if n < len(res) {
		return res[n]
	}

	// init result
	for i := len(res); i <= n; i++ {
		res = append(res, countInt(i))
	}

	return res[n]
}
