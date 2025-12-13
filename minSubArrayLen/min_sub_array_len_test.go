package minSubArrayLen

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "Example 1",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			name:   "Example 2",
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},
		{
			name:   "Example 3",
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSubArrayLen(tt.target, tt.nums)
			if got != tt.want {
				t.Errorf("%v minSubArrayLen() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkMinSubArrayLen(b *testing.B) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	for i := 0; i < b.N; i++ {
		minSubArrayLen(target, nums)
	}
}

func BenchmarkMinSubArrayLenVer2(b *testing.B) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	for i := 0; i < b.N; i++ {
		minSubArrayLenVer2(target, nums)
	}
}
