package main

func lengthOfLastWord(s string) int {
	var res = []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i] >= 'A' && s[i] <= 'Z') ||
			(s[i] >= 'a' && s[i] <= 'z') {
			res = append(res, s[i])
		} else {
			if len(res) > 0 {
				break
			}
		}
	}

	return len(res)
}
