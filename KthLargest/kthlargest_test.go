package KthLargest

import "testing"

func TestKthLargest_Add(t *testing.T) {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	if kthLargest.Add(3) != 4 {
		t.Error("Add 3 failed")
	}
	if kthLargest.Add(5) != 5 {
		t.Error("Add 5 failed")
	}
	if kthLargest.Add(10) != 5 {
		t.Error("Add 10 failed")
	}
	if kthLargest.Add(9) != 8 {
		t.Error("Add 9 failed")
	}
	if kthLargest.Add(4) != 8 {
		t.Error("Add 4 failed")
	}
}
