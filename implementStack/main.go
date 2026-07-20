package main

import "fmt"

// 225. Implement Stack using Queues

// Implement a last-in-first-out (LIFO) stack using only two queues.
// The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).

// Implement the MyStack class:

// void push(int x) Pushes element x to the top of the stack.
// int pop() Removes the element on the top of the stack and returns it.
// int top() Returns the element on the top of the stack.
// boolean empty() Returns true if the stack is empty, false otherwise.
// Notes:

// You must use only standard operations of a queue, which means that only push to back,
//
//	peek/pop from front, size and is empty operations are valid.
//
// Depending on your language, the queue may not be supported natively. You may simulate a queue using
//
//	a list or deque (double-ended queue) as long as you use only a queue's standard operations.
type MyStack struct {
	values []int
}

func Constructor() MyStack {
	return MyStack{
		values: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.values = append(this.values, x)
}

func (this *MyStack) Pop() int {
    lastIdx := len(this.values)-1
    val := this.values[lastIdx]
	this.values = this.values[:lastIdx]
	return val
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.values[len(this.values)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.values) == 0
}

func main() {
	x := 10
	obj := Constructor()
	obj.Push(x)
	param_2 := obj.Pop()
	fmt.Println(param_2)
	param_3 := obj.Top()
	fmt.Println(param_3)
	param_4 := obj.Empty()
	fmt.Println(param_4)
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
