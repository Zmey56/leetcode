package isToeplitzMatrix

import "testing"

func Test_isToeplitzMatrixVer2(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   bool
	}{
		{
			name: "Example 1",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 1, 2, 3},
				{9, 5, 1, 2},
			},
			want: true,
		},
		{
			name: "Example 2",
			matrix: [][]int{
				{1, 2},
				{2, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isToeplitzMatrixVer2(tt.matrix); got != tt.want {
				t.Errorf("isToeplitzMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isToeplitzMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   bool
	}{
		{
			name: "Example 1",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 1, 2, 3},
				{9, 5, 1, 2},
			},
			want: true,
		},
		{
			name: "Example 2",
			matrix: [][]int{
				{1, 2},
				{2, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isToeplitzMatrix(tt.matrix); got != tt.want {
				t.Errorf("isToeplitzMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isToeplitzMatrix(b *testing.B) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}
	for i := 0; i < b.N; i++ {
		isToeplitzMatrix(matrix)
	}
}

func Benchmark_isToeplitzMatrixVer2(b *testing.B) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}
	for i := 0; i < b.N; i++ {
		isToeplitzMatrixVer2(matrix)
	}
}
