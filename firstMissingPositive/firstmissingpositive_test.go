package firstMissingPositive

import "testing"

// Test for function firstMissingPositive

func TestFirstMissingPositive(t *testing.T) {
	test := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test #1",
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			name: "Test #2",
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			name: "Test #3",
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
		{
			name: "Test #4",
			nums: []int{1, 2, 3, 4, 5},
			want: 6,
		},
		{
			name: "Test #5",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 11,
		},
		{
			name: "Test #6",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			want: 12,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}

}
