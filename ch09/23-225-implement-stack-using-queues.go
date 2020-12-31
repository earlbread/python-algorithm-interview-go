package ch09

import "container/list"

type MyStack struct {
	q *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{q: list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q.PushBack(x)
	for i := 0; i < this.q.Len()-1; i++ {
		e := this.q.Front()
		this.q.MoveToBack(e)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	e := this.q.Front()
	this.q.Remove(e)
	return e.Value.(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q.Front().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
