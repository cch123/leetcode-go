package main

func main() {}

func singleNumber(nums []int) []int {
	bitMask := 0
	for _, num := range nums {
		bitMask = bitMask ^ num
	}
	singleBitMask := 1
	for singleBitMask&bitMask == 0 {
		singleBitMask <<= 1
	}
	n1, n2 := 0, 0
	for _, num := range nums {
		if num&singleBitMask == 0 {
			n1 = n1 ^ num
		} else {
			n2 = n2 ^ num
		}
	}
	return []int{n1, n2}
}
