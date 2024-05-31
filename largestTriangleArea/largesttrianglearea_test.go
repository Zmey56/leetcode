package largestTriangleArea

import "testing"

func Test_largestTriangleArea(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  float64
	}{
		{
			name:  "testcase 1",
			input: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
			want:  2.0,
		},
		{
			name:  "testcase 2",
			input: [][]int{{1, 0}, {0, 0}, {0, 1}},
			want:  0.5,
		},
		{
			name:  "testcase 3",
			input: [][]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}},
			want:  0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestTriangleArea(tt.input); got != tt.want {
				t.Errorf("largestTriangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_largestTriangleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}})
	}
}
