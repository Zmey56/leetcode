package minReorder

import "testing"

//test for function minReorder

func TestMinReorder(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		connections [][]int
		want        int
	}{
		{
			name:        "Example 1",
			n:           6,
			connections: [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}},
			want:        3,
		},
		{
			name:        "Example 2",
			n:           5,
			connections: [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}},
			want:        2,
		},
		{
			name:        "Example 3",
			n:           3,
			connections: [][]int{{1, 0}, {2, 0}},
			want:        0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minReorder(tt.n, tt.connections)
			if got != tt.want {
				t.Errorf("%v minReorder() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
