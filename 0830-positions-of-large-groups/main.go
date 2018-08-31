package main

func main() {
}

func largeGroupPositions(S string) [][]int {
	res := [][]int{}
	start := 0
	for i := 1; i < len(S); i++ {
		if S[i] != S[i-1] {
			end := i - 1
			if end-start+1 >= 3 {
				res = append(res, []int{start, end})
			}
			start = i
		}

		// final group
		if i == len(S)-1 && i-start+1 >= 3 {
			res = append(res, []int{start, i})
		}
	}
	return res
}
