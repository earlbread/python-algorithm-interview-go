package ch09

type MyQueue struct {
	input  []int
	output []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		input:  []int{},
		output: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.Peek()

	top := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]
	return last
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			top := this.input[len(this.input)-1]
			this.output = append(this.output, last)
			this.input = this.input[:len(this.input)-1]
		}
	}

	return this.output[len(this.output)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.input) == 0 && len(this.output) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
