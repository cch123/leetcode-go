package main

func main() {
	obj := Constructor()
	obj.Push(4)
	obj.Push(0)
	obj.Push(9)
	obj.Push(3)
	obj.Push(4)
	obj.Push(2)
	obj.Pop()
	obj.Push(6)
	obj.Pop()
	obj.Push(1)
	obj.Pop()
	obj.Push(1)
	obj.Pop()
	obj.Push(4)
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
	obj.Pop()
}

type FreqStack struct {
	freq       map[int]int
	freq2stack map[int][]int
	maxFreq    int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:       make(map[int]int),
		freq2stack: map[int][]int{},
		maxFreq:    0,
	}
}

func (this *FreqStack) Push(x int) {
	this.freq[x]++
	if this.freq[x] > this.maxFreq {
		this.maxFreq = this.freq[x]
	}
	this.freq2stack[this.freq[x]] = append(this.freq2stack[this.freq[x]], x)
}

func (this *FreqStack) Pop() int {
	lastIdx := len(this.freq2stack[this.maxFreq]) - 1
	res := this.freq2stack[this.maxFreq][lastIdx]
	this.freq2stack[this.maxFreq] = this.freq2stack[this.maxFreq][:lastIdx] // pop
	this.freq[res]--

	if len(this.freq2stack[this.maxFreq]) == 0 {
		delete(this.freq2stack, this.maxFreq)
		this.maxFreq--
	}
	return res
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */
