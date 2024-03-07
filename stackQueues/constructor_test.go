package stackQueues

import "testing"

//Test for MyStack

func TestMyStack(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	tmpTop := obj.Top()
	if tmpTop != 2 {
		t.Errorf("obj.Top() = %d; want 2", tmpTop)
	}
	tmpPop := obj.Pop()
	if tmpPop != 2 {
		t.Errorf("obj.Pop() = %d; want 2", tmpPop)
	}
}
