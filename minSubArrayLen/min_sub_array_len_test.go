package minSubArrayLen

import (
	"math/rand"
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
			got := minSubArrayLenV3(tt.target, tt.nums)
			if got != tt.want {
				t.Errorf("%v minSubArrayLen() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkMinSubArrayLen(b *testing.B) {
	testCases := []struct {
		name   string
		target int
		size   int
	}{
		{"Small_100", 500, 100},
		{"Medium_1K", 5000, 1000},
		{"Large_10K", 50000, 10000},
		{"VeryLarge_50K", 250000, 50000},
		{"Huge_100K", 500000, 100000},
	}

	for _, tt := range testCases {
		nums := generateArray(tt.size)

		b.Run("BenchmarkMinSubArrayLen", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				minSubArrayLen(tt.target, nums)
			}
		})

		b.Run("BenchmarkMinSubArrayLenParallel", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				minSubArrayLenVer2(tt.target, nums)
			}
		})

		b.Run("BenchmarkMinSubArrayLenParallel", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				minSubArrayLenV3(tt.target, nums)
			}
		})

		b.Run("BenchmarkMinSubArrayLenParallel", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				minSubArrayLenVer4(tt.target, nums)
			}
		})
	}

}

func generateArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100) + 1
	}
	return arr
}
