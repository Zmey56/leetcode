package maxCount

import "testing"

// Test for function maxCount

func Test_maxCount(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		ops  [][]int
		want int
	}{
		{"Test 1", 3, 3, [][]int{{2, 2}, {3, 3}}, 4},
		{"Test 2", 3, 3, [][]int{{2, 2}, {3, 3}, {1, 1}}, 1},
		{"Test 3", 3, 3, [][]int{}, 9},
		{"Test 4", 3, 3, [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}}, 4},
	}

	for _, tt := range tests {
		if got := maxCount(tt.m, tt.n, tt.ops); got != tt.want {
			t.Errorf("%v maxCount(%v, %v, %v) = %v; want %v", tt.name, tt.m, tt.n, tt.ops, got, tt.want)
		}
	}
}
