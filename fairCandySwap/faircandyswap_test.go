package fairCandySwap

import (
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		B    []int
		want []int
	}{
		{
			name: "Example 1",
			A:    []int{1, 1},
			B:    []int{2, 2},
			want: []int{1, 2},
		},
		{
			name: "Example 2",
			A:    []int{1, 2},
			B:    []int{2, 3},
			want: []int{1, 2},
		},
		{
			name: "Example 3",
			A:    []int{2},
			B:    []int{1, 3},
			want: []int{2, 3},
		},
		{
			name: "Example 4",
			A:    []int{1, 2, 5},
			B:    []int{2, 4},
			want: []int{5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwap(tt.A, tt.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
