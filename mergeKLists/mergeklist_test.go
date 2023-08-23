package main

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{
		&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 2, Next: &ListNode{Val: 6}},
	}

	expected := []int{1, 1, 2, 3, 4, 4, 5, 6}

	result := mergeKLists(lists)

	for i := 0; i < len(expected); i++ {
		if result == nil {
			t.Errorf("Result list is shorter than expected.")
			break
		}
		if result.Val != expected[i] {
			t.Errorf("Expected %d, but got %d", expected[i], result.Val)
		}
		result = result.Next
	}

	if result != nil {
		t.Errorf("Result list is longer than expected.")
	}
}
