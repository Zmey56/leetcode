package distributeCandies

import "testing"

// Test for function distributeCandies
func Test_distributeCandies(t *testing.T) {
	tests := []struct {
		name      string
		candyType []int
		want      int
	}{
		{"Test 1", []int{1, 1, 2, 2, 3, 3}, 3},
		{"Test 2", []int{1, 1, 2, 3}, 2},
		{"Test 3", []int{6, 6, 6, 6}, 1},
	}

	for _, tt := range tests {
		if got := distributeCandies(tt.candyType); got != tt.want {
			t.Errorf("%v distributeCandies(%v) = %v; want %v", tt.name, tt.candyType, got, tt.want)
		}
	}
}
