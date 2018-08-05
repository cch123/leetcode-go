package main

func main() {

}

func sliceEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var byteMapS = map[byte][]int{}
	var byteMapT = map[byte][]int{}
	for i := 0; i < len(s); i++ {
		currentChar := s[i]
		byteMapS[currentChar] = append(byteMapS[currentChar], i)
		currentCharT := t[i]
		byteMapT[currentCharT] = append(byteMapT[currentCharT], i)
		if !sliceEqual(byteMapS[currentChar], byteMapT[currentCharT]) {
			return false
		}
	}
	return true
}
