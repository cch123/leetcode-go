package main

import "fmt"

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
}

func getWordMap(A string) map[string]int {
	mapA := map[string]int{}
	currentWord := []byte{}
	for i := 0; i < len(A); i++ {
		if A[i] >= 'a' && A[i] <= 'z' {
			currentWord = append(currentWord, A[i])
		} else {
			mapA[string(currentWord)]++
			currentWord = currentWord[:0]
		}
	}
	println(string(currentWord))
	if len(currentWord) > 0 {
		mapA[string(currentWord)]++
	}
	return mapA
}

func uncommonFromSentences(A string, B string) []string {
	var mapA = getWordMap(A)
	var mapB = getWordMap(B)
	fmt.Println(mapA, mapB)
	res := []string{}
	for k, cnt := range mapA {
		if mapB[k] == 0 && cnt == 1 {
			res = append(res, k)
		}
	}
	for k, cnt := range mapB {
		if mapA[k] == 0 && cnt == 1 {
			res = append(res, k)
		}
	}
	return res
}
