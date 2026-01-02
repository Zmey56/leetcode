package containsNearbyDuplicateV2

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	test := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, tt := range test {
		got := containsNearbyDuplicate(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("containsNearbyDuplicate(%v, %v) = %v; want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
