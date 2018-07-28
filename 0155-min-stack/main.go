package main

type MinStack struct {
	elems []int
	mins  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		elems: make([]int, 0),
		mins:  make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.elems = append(this.elems, x)
	topIdx := len(this.mins) - 1
	if topIdx >= 0 && this.mins[topIdx] < x {
		this.mins = append(this.mins, this.mins[topIdx])
	} else {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop() {
	l := len(this.elems)
	this.elems = this.elems[0 : l-1]
	this.mins = this.mins[0 : l-1]
}

func (this *MinStack) Top() int {
	topIdx := len(this.elems) - 1
	return this.elems[topIdx]
}

func (this *MinStack) GetMin() int {
	topIdx := len(this.mins) - 1
	return this.mins[topIdx]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
