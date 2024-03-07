package stackQueues

import "log"

//Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all
//the functions of a normal stack (push, top, pop, and empty).
//
//Implement the MyStack class:
//
//void push(int x) Pushes element x to the top of the stack.
//int pop() Removes the element on the top of the stack and returns it.
//int top() Returns the element on the top of the stack.
//boolean empty() Returns true if the stack is empty, false otherwise.

type MyStack struct {
	Queue []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	log.Println("PUSH", this.Queue)
	this.Queue = append(this.Queue, x)
}

func (this *MyStack) Pop() int {
	log.Println("POP", this.Queue, len(this.Queue), this.Queue[len(this.Queue)-1])
	result := this.Queue[len(this.Queue)-1]
	this.Queue = this.Queue[:len(this.Queue)-1]
	log.Println("POP 2", this.Queue)
	return result
}

func (this *MyStack) Top() int {
	log.Println("TOP", this.Queue)
	return this.Queue[len(this.Queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.Queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
