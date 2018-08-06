package main

import "strings"

func main() {}

func countSegments(s string) int {
	arr := strings.Split(s, " ")
	cnt := 0
	for _, x := range arr {
		if len(x) > 0 {
			cnt++
		}
	}
	return cnt
}
