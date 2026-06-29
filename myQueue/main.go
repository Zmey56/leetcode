package main

type MyQueue struct {
	store []int
}

func Constructor() MyQueue {
	return MyQueue{
		store: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.store = append(this.store, x)
}

func (this *MyQueue) Pop() int {
	value := this.store[0]
	this.store = this.store[1:]
	return value
}

func (this *MyQueue) Peek() int {
	return this.store[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.store) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
