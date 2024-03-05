package containsNearbyDuplicate

import "testing"

// Test containsNearbyDuplicate function

func TestContainsNearbyDuplicate(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, test := range tests {
		if got := containsNearbyDuplicate(test.nums, test.k); got != test.want {
			t.Errorf("containsNearbyDuplicate(%v, %v) = %v; want %v", test.nums, test.k, got, test.want)
		}
	}
}
