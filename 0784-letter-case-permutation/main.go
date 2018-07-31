package main

import (
	"fmt"
	"strings"
)

func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return []string{S}
	}

	if len(S) == 1 {
		if S[0] >= '0' && S[0] <= '9' {
			return []string{S}
		}

		// 说明是字母
		return []string{strings.ToLower(S), strings.ToUpper(S)}
	}

	subStringResult := letterCasePermutation(S[1:])

	result := []string{}
	var leadingStrings []string
	if S[0] >= '0' && S[0] <= '9' {
		leadingStrings = []string{S[:1]}
	} else {
		leadingStrings = []string{
			strings.ToLower(S[:1]),
			strings.ToUpper(S[:1]),
		}
	}

	for i := 0; i < len(leadingStrings); i++ {
		for j := 0; j < len(subStringResult); j++ {
			result = append(result, leadingStrings[i]+subStringResult[j])
		}
	}

	return result
}

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("12345"))
	fmt.Println(letterCasePermutation("3z4"))
}
