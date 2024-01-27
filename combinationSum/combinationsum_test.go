package combinationSum

import "testing"

// Create TestCombinationSum

func TestCombinationSum(t *testing.T) {
	test := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "Test #1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want: [][]int{
				[]int{2, 2, 3},
				[]int{7},
			},
		},
		{
			name:       "Test #2",
			candidates: []int{2, 3, 5},
			target:     8,
			want: [][]int{
				[]int{2, 2, 2, 2},
				[]int{2, 3, 3},
				[]int{3, 5},
			},
		},
		{
			name:       "Test #3",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
		{
			name:       "Test #4",
			candidates: []int{1},
			target:     1,
			want: [][]int{
				[]int{1},
			},
		},
		{
			name:       "Test #5",
			candidates: []int{1},
			target:     2,
			want: [][]int{
				[]int{1, 1},
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.candidates, tt.target); !compare(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// compare compares two 2D slices of ints
func compare(a, b [][]int) bool {
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
