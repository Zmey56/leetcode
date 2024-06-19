package minIncrementForUnique

import "testing"

func Test_minIncrementForUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "testcase 1",
			nums: []int{1, 2, 2},
			want: 1,
		},
		{
			name: "testcase 2",
			nums: []int{3, 2, 1, 2, 1, 7},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minIncrementForUnique(tt.nums); got != tt.want {
				t.Errorf("minIncrementForUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_minIncrementForUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minIncrementForUnique([]int{3, 2, 1, 2, 1, 7})
	}
}

func Benchmark_minIncrementForUniqueVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minIncrementForUniqueVer2([]int{3, 2, 1, 2, 1, 7})
	}
}
