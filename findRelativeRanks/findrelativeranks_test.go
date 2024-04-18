package findRelativeRanks

import "testing"

//Test for function findRelativeRanks

func Test_findRelativeRanks(t *testing.T) {
	type args struct {
		score []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Test case 1",
			args{[]int{10, 3, 8, 9, 4}},
			[]string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
		{
			"Test case 2",
			args{[]int{5, 4, 3, 2, 1}},
			[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			"Test case 3",
			args{[]int{1}},
			[]string{"Gold Medal"},
		},
		{
			"Test case 4",
			args{[]int{1, 2}},
			[]string{"Silver Medal", "Gold Medal"},
		},
		{
			"Test case 5",
			args{[]int{1, 2, 3}},
			[]string{"Bronze Medal", "Silver Medal", "Gold Medal"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRelativeRanks(tt.args.score); !compareStringSlices(got, tt.want) {
				t.Errorf("findRelativeRanks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
