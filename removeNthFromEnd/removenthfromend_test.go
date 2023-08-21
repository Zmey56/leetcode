package main

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	t.Run("TestCases", func(t *testing.T) {
		testCases := []struct {
			head     []int
			n        int
			expected []int
		}{
			{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
			{[]int{1}, 1, []int{}},
			{[]int{1, 2}, 1, []int{1}},
			{[]int{1, 2, 3}, 3, []int{2, 3}},
			{[]int{1, 2, 3}, 1, []int{1, 2}},
			// Add more test cases as needed
		}

		for _, tc := range testCases {
			t.Run("", func(t *testing.T) {
				head := createLinkedList(tc.head)
				newHead := removeNthFromEnd(head, tc.n)
				actual := displayLinkedList(newHead)
				if !reflect.DeepEqual(actual, tc.expected) {
					t.Errorf("Input: %v, n: %d\nExpected: %v\nActual: %v", tc.head, tc.n, tc.expected, actual)
				}
			})
		}
	})
}
