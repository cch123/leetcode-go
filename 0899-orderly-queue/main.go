package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(orderlyQueue("cba", 1))
	fmt.Println(orderlyQueue("baaca", 3))
}

func orderlyQueue(S string, K int) string {
	if K >= 2 {
		sBytes := []byte(S)
		sort.Slice(sBytes, func(i, j int) bool {
			return sBytes[i] < sBytes[j]
		})
		return string(sBytes)
	}
	minStr := S
	for i := 1; i < len(S); i++ {
		curStr := S[i:] + S[0:i]
		if curStr < minStr {
			minStr = curStr
		}
	}
	return minStr
}
