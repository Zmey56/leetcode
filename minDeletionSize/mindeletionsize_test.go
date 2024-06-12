package minDeletionSize

import "testing"

func Test_minDeletionSize(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want int
	}{
		{
			name: "Example 1",
			strs: []string{"cba", "daf", "ghi"},
			want: 1,
		},
		{
			name: "Example 2",
			strs: []string{"a", "b"},
			want: 0,
		},
		{
			name: "Example 3",
			strs: []string{"zyx", "wvu", "tsr"},
			want: 3,
		},
	}
	for _, tt := range tests {
		if got := minDeletionSize(tt.strs); got != tt.want {
			t.Errorf("%q. minDeletionSize() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_minDeletionSizeVer2(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want int
	}{
		{
			name: "Example 1",
			strs: []string{"cba", "daf", "ghi"},
			want: 1,
		},
		{
			name: "Example 2",
			strs: []string{"a", "b"},
			want: 0,
		},
		{
			name: "Example 3",
			strs: []string{"zyx", "wvu", "tsr"},
			want: 3,
		},
	}
	for _, tt := range tests {
		if got := minDeletionSize2(tt.strs); got != tt.want {
			t.Errorf("%q. minDeletionSize() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_minDeletionSizeVer3(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want int
	}{
		{
			name: "Example 1",
			strs: []string{"cba", "daf", "ghi"},
			want: 1,
		},
		{
			name: "Example 2",
			strs: []string{"a", "b"},
			want: 0,
		},
		{
			name: "Example 3",
			strs: []string{"zyx", "wvu", "tsr"},
			want: 3,
		},
	}
	for _, tt := range tests {
		if got := minDeletionSize3(tt.strs); got != tt.want {
			t.Errorf("%q. minDeletionSize() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Benchmark_minDeletionSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minDeletionSize([]string{"cba", "daf", "ghi"})
	}
}

func Benchmark_minDeletionSizeVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minDeletionSize2([]string{"cba", "daf", "ghi"})
	}
}

func Benchmark_minDeletionSizeVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minDeletionSize3([]string{"cba", "daf", "ghi"})
	}
}
