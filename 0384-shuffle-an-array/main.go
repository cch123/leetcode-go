package main

import "math/rand"

func main() {

}

type Solution struct {
	initial []int
}

func Constructor(nums []int) Solution {
	initial := append([]int{}, nums...)
	return Solution{initial: initial}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.initial
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	copied := append([]int{}, this.initial...)
	for i := 0; i < len(this.initial); i++ {
		elemCnt := len(this.initial) - i
		j := rand.Intn(elemCnt)
		copied[j], copied[elemCnt-1] = copied[elemCnt-1], copied[j]
	}
	return copied
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
