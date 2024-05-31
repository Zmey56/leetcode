package lemonadeChange

import "testing"

func Test_lemonadeChange(t *testing.T) {
	tests := []struct {
		name  string
		bills []int
		want  bool
	}{
		{
			name:  "testcase 1",
			bills: []int{5, 5, 5, 10, 20},
			want:  true,
		},
		{
			name:  "testcase 2",
			bills: []int{5, 5, 10},
			want:  true,
		},
		{
			name:  "testcase 3",
			bills: []int{10, 10},
			want:  false,
		},
		{
			name:  "testcase 4",
			bills: []int{5, 5, 10, 10, 20},
			want:  false,
		},
		{
			name:  "testcase 5",
			bills: []int{5, 5, 10, 10, 10},
			want:  false,
		},
		{
			name:  "testcase 6",
			bills: []int{5, 5, 5, 5, 20, 20, 5, 5, 5, 5},
			want:  false,
		},
		{
			name:  "testcase 7",
			bills: []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 20},
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lemonadeChange(tt.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
