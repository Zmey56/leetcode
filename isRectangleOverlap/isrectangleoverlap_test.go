package isRectangleOverlap

import "testing"

func Test_isRectangleOverlap(t *testing.T) {
	tests := []struct {
		name string
		rec1 []int
		rec2 []int
		want bool
	}{
		{
			name: "testcase 1",
			rec1: []int{0, 0, 2, 2},
			rec2: []int{1, 1, 3, 3},
			want: true,
		},
		{
			name: "testcase 2",
			rec1: []int{0, 0, 1, 1},
			rec2: []int{1, 0, 2, 1},
			want: false,
		},
		{
			name: "testcase 3",
			rec1: []int{0, 0, 1, 1},
			rec2: []int{2, 2, 3, 3},
			want: false,
		},
		{
			name: "testcase 4",
			rec1: []int{7, 8, 13, 15},
			rec2: []int{10, 8, 12, 20},
			want: true,
		},
		{
			name: "testcase 5",
			rec1: []int{0, 0, 1, 1},
			rec2: []int{1, 0, 2, 1},
			want: false,
		},
		{
			name: "testcase 6",
			rec1: []int{0, 0, 1, 1},
			rec2: []int{1, 1, 2, 2},
			want: false,
		},
		{
			name: "testcase 7",
			rec1: []int{0, 0, 1, 1},
			rec2: []int{0, 0, 1, 1},
			want: true,
		},
		{
			name: "testcase 8",
			rec1: []int{5, 15, 8, 18},
			rec2: []int{0, 3, 7, 9},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleOverlap(tt.rec1, tt.rec2); got != tt.want {
				t.Errorf("isRectangleOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
