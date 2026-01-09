package MinStack

// 155. Min Stack

//Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
//Implement the MinStack class:
//
//MinStack() initializes the stack object.
//void push(int val) pushes the element val onto the stack.
//void pop() removes the element on the top of the stack.
//int top() gets the top element of the stack.
//int getMin() retrieves the minimum element in the stack.
//You must implement a solution with O(1) time complexity for each function.

type Node struct {
	val int
	min int
}

type MinStack struct {
	stack []Node
}

func Constructor() MinStack {
	return MinStack{
		stack: []Node{},
	}

}

func (this *MinStack) Push(val int) {
	newMin := val
	if len(this.stack) > 0 {
		currentMin := this.stack[len(this.stack)-1].min
		if currentMin < newMin {
			newMin = currentMin
		}
	}
	this.stack = append(this.stack, Node{val, newMin})
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
