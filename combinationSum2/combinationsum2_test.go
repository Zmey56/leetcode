package combinationSum2

import "testing"

//create test for function combinationSum2

func TestCombinationSum2(t *testing.T) {
	test := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "Test #1",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			want: [][]int{
				[]int{1, 1, 6},
				[]int{1, 2, 5},
				[]int{1, 7},
				[]int{2, 6},
			},
		},
		{
			name:       "Test #2",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			want: [][]int{
				[]int{1, 2, 2},
				[]int{5},
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2(tt.candidates, tt.target); !compareTest(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v, target = %v", got, tt.want, tt.target)
			}
		})
	}

}

// compare compares two 2D slices of ints
func compareTest(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
