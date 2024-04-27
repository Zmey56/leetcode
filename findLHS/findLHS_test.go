package findLHS

import "testing"

// Test for function findLHS
func Test_findLHS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Test 1", []int{1, 3, 2, 2, 5, 2, 3, 7}, 5},
		{"Test 2", []int{1, 2, 3, 4}, 2},
		{"Test 3", []int{1, 1, 1, 1}, 0},
	}

	for _, tt := range tests {
		if got := findLHS(tt.nums); got != tt.want {
			t.Errorf("%v findLHS(%v) = %v; want %v", tt.name, tt.nums, got, tt.want)
		}
	}
}
