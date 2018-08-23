package main

func main() {}

func canConstruct(ransomNote string, magazine string) bool {
	byteMap := map[byte]int{}
	for i := 0; i < len(magazine); i++ {
		byteMap[magazine[i]]++
	}
	byteMapForR := map[byte]int{}
	for i := 0; i < len(ransomNote); i++ {
		byteMapForR[ransomNote[i]]++
		if byteMap[ransomNote[i]] < byteMapForR[ransomNote[i]] {
			return false
		}
	}
	return true

}
