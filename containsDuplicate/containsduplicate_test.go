package containsDuplicate

import "testing"

//Test for function containsDuplicate

func TestContainsDuplicate(t *testing.T) {
	var tests = []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, test := range tests {
		if got := containsDuplicate(test.nums); got != test.want {
			t.Errorf("containsDuplicate(%v) = %v", test.nums, got)
		}
	}
}
