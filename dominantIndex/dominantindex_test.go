package dominantIndex

import "testing"

func Test_dominantIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Example 1", []int{3, 6, 1, 0}, 1},
		{"Example 2", []int{1, 2, 3, 4}, -1},
		{"Example 3", []int{1}, 0},
		{"Example 4", []int{1, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dominantIndex(tt.nums); got != tt.want {
				t.Errorf("dominantIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
