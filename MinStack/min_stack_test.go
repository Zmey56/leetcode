package MinStack

import "testing"

func TestMinStack(t *testing.T) {
	minStack := Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	// Проверка GetMin
	if res := minStack.GetMin(); res != -3 {
		t.Errorf("Expected GetMin to be -3, got %d", res)
	}

	minStack.Pop()

	// Проверка Top
	if res := minStack.Top(); res != 0 {
		t.Errorf("Expected Top to be 0, got %d", res)
	}

	// Проверка GetMin после Pop
	if res := minStack.GetMin(); res != -2 {
		t.Errorf("Expected GetMin to be -2, got %d", res)
	}
}
