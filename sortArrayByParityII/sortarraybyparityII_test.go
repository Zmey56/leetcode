package sortArrayByParityII

import (
	"sort"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Test case 1",
			nums: []int{4, 2, 5, 7},
			want: []int{4, 5, 2, 7},
		},
		{
			name: "Test case 2",
			nums: []int{2, 3},
			want: []int{2, 3},
		},
		{
			name: "Test case 3",
			nums: []int{3, 2},
			want: []int{2, 3},
		},
		{
			name: "Test case 4",
			nums: []int{3, 2, 4, 1},
			want: []int{2, 3, 4, 1},
		},
		{
			name: "Test case 5",
			nums: []int{1, 2, 3, 4},
			want: []int{2, 1, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParityII(tt.nums); !deepEqualIgnoreOrder(got, tt.want) {
				t.Errorf("sortArrayByParityII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func deepEqualIgnoreOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
