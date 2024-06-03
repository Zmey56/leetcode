package sortArrayByParity

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Test case 1",
			nums: []int{3, 1, 2, 4},
			want: []int{2, 4, 3, 1},
		},
		{
			name: "Test case 2",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "Test case 3",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "Test case 4",
			nums: []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "Test case 5",
			nums: []int{2, 1},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArrayByParity(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_sortArrayByParity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArrayByParity([]int{3, 1, 2, 4})
	}
}

func Benchmark_sortArrayByParityVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArrayByParityVer2([]int{3, 1, 2, 4})
	}
}
