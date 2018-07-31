package main

func numJewelsInStones(J string, S string) int {
	var diamondMap = map[byte]bool{}
	for i := 0; i < len(J); i++ {
		diamondMap[J[i]] = true
	}

	cnt := 0
	for i := 0; i < len(S); i++ {
		if diamondMap[S[i]] == true {
			cnt++
		}
	}
	return cnt
}
