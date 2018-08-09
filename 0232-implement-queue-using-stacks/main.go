package main

func main() {

}

type MyQueue struct {
	stackLeft  []int
	stackRight []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stackLeft:  []int{},
		stackRight: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if len(this.stackLeft) == 0 && len(this.stackRight) != 0 {
		for i := len(this.stackRight) - 1; i >= 0; i-- {
			this.stackLeft = append(this.stackLeft, this.stackRight[i])
		}
		this.stackRight = this.stackRight[:0]
	}

	this.stackLeft = append(this.stackLeft, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.stackLeft) != 0 {
		for i := len(this.stackLeft) - 1; i >= 0; i-- {
			this.stackRight = append(this.stackRight, this.stackLeft[i])
		}
		this.stackLeft = this.stackLeft[:0]
	}
	popElem := this.stackRight[len(this.stackRight)-1]
	this.stackRight = this.stackRight[:len(this.stackRight)-1]
	return popElem
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stackLeft) > 0 {
		return this.stackLeft[0]
	}
	return this.stackRight[len(this.stackRight)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.stackRight) > 0 || len(this.stackLeft) > 0 {
		return false
	}

	return true
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
