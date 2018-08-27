package main

import "fmt"

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "adcb", "cbad"}))
	fmt.Println(numSpecialEquivGroups([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
}

func isInSameGroup(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	oddMapForA := map[byte]int{}
	evenMapForA := map[byte]int{}
	oddMapForB := map[byte]int{}
	evenMapForB := map[byte]int{}
	for i := 0; i < len(a); i++ {
		if i%2 == 0 {
			oddMapForA[a[i]]++
			oddMapForB[b[i]]++
			continue
		}
		evenMapForA[a[i]]++
		evenMapForB[b[i]]++
	}
	for b, cnt := range oddMapForA {
		if cnt != oddMapForB[b] {
			return false
		}
	}
	for b, cnt := range evenMapForA {
		if cnt != evenMapForB[b] {
			return false
		}
	}
	return true
}

func numSpecialEquivGroups(A []string) int {
	used := make([]bool, len(A))
	cnt := 0
	for i := 0; i < len(A); i++ {
		if used[i] == true {
			continue
		}
		used[i] = true
		cnt++

		group := []string{A[i]}
		for j := i + 1; j < len(A); j++ {
			if isInSameGroup(A[j], A[i]) {
				used[j] = true
				group = append(group, A[j])
			}
		}
	}
	return cnt
}
