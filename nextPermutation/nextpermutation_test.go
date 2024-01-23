package nextPermutation

import "testing"

//Create test for function nextPermutation

func TestNextPermutation(t *testing.T) {
	test := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
		{
			nums: []int{1},
			want: []int{1},
		},
	}

	for _, test := range test {
		nextPermutation(test.nums)
		if !equal(test.nums, test.want) {
			t.Errorf("nextPermutation(%v) = %v, want %v", test.nums, test.nums, test.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true

}
