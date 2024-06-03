package isMonotonic

import "testing"

func Test_isMonotonic(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Test case 1",
			nums: []int{1, 2, 2, 3},
			want: true,
		},
		{
			name: "Test case 2",
			nums: []int{6, 5, 4, 4},
			want: true,
		},
		{
			name: "Test case 3",
			nums: []int{1, 3, 2},
			want: false,
		},
		{
			name: "Test case 4",
			nums: []int{1, 2, 4, 5},
			want: true,
		},
		{
			name: "Test case 5",
			nums: []int{1, 1, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.nums); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
