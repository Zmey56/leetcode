package majorityElement

import "testing"

//Test for function majorityElement

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
	}

	for _, test := range tests {
		if got := majorityElement(test.nums); got != test.want {
			t.Errorf("majorityElement(%v) = %v; want %v", test.nums, got, test.want)
		}
	}
}
