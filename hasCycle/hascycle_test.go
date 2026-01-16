package hasCycle

import "testing"

// Test for hasCycle function
func TestHasCycle(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		pos      int
		expected bool
	}{
		{
			name:     "Test 1: Cycle present",
			input:    []int{3, 2, 0, -4},
			pos:      1,
			expected: true,
		},
		{
			name:     "Test 2: Cycle present",
			input:    []int{1, 2},
			pos:      0,
			expected: true,
		},
		{
			name:     "Test 3: No cycle",
			input:    []int{1},
			pos:      -1,
			expected: false,
		},
		{
			name:     "Test 4: No cycle",
			input:    []int{1, 2, 3, 4, 5},
			pos:      -1,
			expected: false,
		},
		{
			name:     "Test 5: Empty list",
			input:    []int{},
			pos:      -1,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			head := createListWithCycle(test.input, test.pos)
			if result := hasCycle(head); result != test.expected {
				t.Errorf("hasCycle failed, expected %v, got %v", test.expected, result)
			}
			if result := hasCycleV2(head); result != test.expected {
				t.Errorf("hasCycleV2 failed, expected %v, got %v", test.expected, result)
			}
			if result := hasCycleWithMap(head); result != test.expected {
				t.Errorf("hasCycleWithMap failed, expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Create a linked list from a slice with a cycle at a given position
func createListWithCycle(arr []int, pos int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	nodes := []*ListNode{head}
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
		nodes = append(nodes, current)
	}

	if pos != -1 && pos < len(nodes) {
		current.Next = nodes[pos]
	}

	return head
}

func generateList(n int) []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i
	}
	return list
}

func BenchmarkHasCycle(b *testing.B) {
	list := createListWithCycle(generateList(1000), 500)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasCycle(list)
	}
}

func BenchmarkHasCycleV2(b *testing.B) {
	list := createListWithCycle(generateList(1000), 500)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasCycleV2(list)
	}
}

func BenchmarkHasCycleWithMap(b *testing.B) {
	list := createListWithCycle(generateList(1000), 500)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasCycleWithMap(list)
	}
}

func BenchmarkHasCycle_NoCycle(b *testing.B) {
	list := createListWithCycle(generateList(1000), -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasCycle(list)
	}
}

func BenchmarkHasCycleV2_NoCycle(b *testing.B) {
	list := createListWithCycle(generateList(1000), -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasCycleV2(list)
	}
}

func BenchmarkHasCycleWithMap_NoCycle(b *testing.B) {
	list := createListWithCycle(generateList(1000), -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasCycleWithMap(list)
	}
}
