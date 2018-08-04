package main

func backspaceCompare(S string, T string) bool {
	sResultBytes := []byte{}
	tResultBytes := []byte{}
	for i := 0; i < len(S); i++ {
		if S[i] != '#' {
			sResultBytes = append(sResultBytes, S[i])
			continue
		}
		if len(sResultBytes) > 0 {
			sResultBytes = sResultBytes[:len(sResultBytes)-1]
		}
	}
	for i := 0; i < len(T); i++ {
		if T[i] != '#' {
			tResultBytes = append(tResultBytes, T[i])
			continue
		}
		if len(tResultBytes) > 0 {
			tResultBytes = tResultBytes[:len(tResultBytes)-1]
		}
	}

	return string(sResultBytes) == string(tResultBytes)
}

func main() {

}
