package main

func main() {}

func reverseVowels(s string) string {
	vowelArr := []byte{}
	var m = map[byte]bool{
		'a': true,
		'A': true,
		'e': true,
		'E': true,
		'i': true,
		'I': true,
		'o': true,
		'u': true,
		'O': true,
		'U': true,
	}

	for i := 0; i < len(s); i++ {
		if m[s[i]] == true {
			vowelArr = append(vowelArr, s[i])
		}
	}

	res := []byte{}
	pos := len(vowelArr) - 1
	for i := 0; i < len(s); i++ {
		if m[s[i]] == true {
			vowelArr = append(vowelArr, s[i])
			res = append(res, vowelArr[pos])
			pos--
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
