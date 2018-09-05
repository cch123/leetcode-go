package main

import "fmt"

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))
}

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 1 {
		return true
	}
	for i := 0; i < len(bits)-1; {
		if bits[i] == 0 {
			i++
			if i == len(bits)-1 {
				return true
			}
			continue
		}
		if bits[i] == 1 && bits[i+1] == 0 {
			i += 2
			if i == len(bits)-1 {
				return true
			}
			continue
		}
		if bits[i] == 1 && bits[i+1] == 1 {
			i += 2
			if i == len(bits)-1 {
				return true
			}
			continue
		}
	}
	return false
}
