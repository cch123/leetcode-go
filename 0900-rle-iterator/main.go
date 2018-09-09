package main

import "fmt"

func main() {
	rie := Constructor([]int{811, 903, 310, 730, 899, 684, 472, 100, 434, 611})
	fmt.Println(rie.Next(358))
	fmt.Println(rie.seq)
	fmt.Println(rie.Next(345))
	fmt.Println(rie.seq)
	fmt.Println(rie.Next(154))
	fmt.Println(rie.seq)
	fmt.Println(rie.Next(265))
	fmt.Println(rie.seq)
	fmt.Println(rie.Next(73))
	fmt.Println(rie.seq)
	fmt.Println(rie.Next(220))
	fmt.Println(rie.seq)
	fmt.Println(rie.Next(138))
	fmt.Println(rie.seq)
	fmt.Println(rie.Next(4))
	fmt.Println(rie.seq)
	fmt.Println(rie.Next(170))
	fmt.Println(rie.seq)
	fmt.Println(rie.Next(88))
	fmt.Println(rie.seq)
}

type RLEIterator struct {
	seq        []int
	currentPos int
}

func Constructor(A []int) RLEIterator {
	rieIter := RLEIterator{
		seq:        A,
		currentPos: 0,
	}

	return rieIter
}

func (this *RLEIterator) Next(n int) int {
	if len(this.seq) == 0 {
		return -1
	}

	pos := this.currentPos
	for {
		if pos >= len(this.seq) {
			break
		}

		if this.seq[pos] >= n {
			this.seq[pos] -= n
			return this.seq[pos+1]
		}

		n = n - this.seq[pos]
		this.seq[pos] = 0
		pos += 2
	}
	return -1
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(A);
 * param_1 := obj.Next(n);
 */
