package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
func subdomainVisits(cpdomains []string) []string {
	resultMap := map[string]int64{}
	for i := 0; i < len(cpdomains); i++ {
		numDomain := strings.Split(cpdomains[i], " ")
		num, _ := strconv.ParseInt(numDomain[0], 10, 64)
		resultMap[numDomain[1]] += num

		// 逆方向遍历 cpdomains[i] 字符串，一旦在第 i 个位置遇到了 .
		// 那么就对 str[i+1:] 进行计数
		for j := len(cpdomains[i]) - 1; j >= 0; j-- {
			if cpdomains[i][j] == '.' {
				subdomain := cpdomains[i][j+1:]
				resultMap[subdomain] += num
			}
		}
	}

	res := []string{}
	for k, v := range resultMap {
		str := fmt.Sprint(v) + " " + k
		res = append(res, str)
	}
	return res
}
