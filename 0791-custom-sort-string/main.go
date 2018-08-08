package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(customSortString("kqep", "pekeq"))
}
func customSortString(S string, T string) string {
	var byteMap = map[byte]int{}
	for i := 0; i < len(S); i++ {
		byteMap[S[i]] = i
	}

	tBytes := []byte(T)

	sort.Slice(tBytes, func(x, y int) bool {
		c1, c2 := tBytes[x], tBytes[y]
		return byteMap[c1] < byteMap[c2]
	})

	return string(tBytes)
}
