package validMountainArray

import "testing"

func Test_validMountainArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "Test case 1",
			arr:  []int{0, 3, 2, 1},
			want: true,
		},
		{
			name: "Test case 2",
			arr:  []int{3, 5, 5},
			want: false,
		},
		{
			name: "Test case 3",
			arr:  []int{0, 3, 2, 1},
			want: true,
		},
		{
			name: "Test case 4",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: false,
		},
		{
			name: "Test case 5",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want: false,
		},
		{
			name: "Test case 6",
			arr:  []int{0, 3, 2, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_validMountainArrayVer2(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "Test case 1",
			arr:  []int{0, 3, 2, 1},
			want: true,
		},
		{
			name: "Test case 2",
			arr:  []int{3, 5, 5},
			want: false,
		},
		{
			name: "Test case 3",
			arr:  []int{0, 3, 2, 1},
			want: true,
		},
		{
			name: "Test case 4",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: false,
		},
		{
			name: "Test case 5",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want: false,
		},
		{
			name: "Test case 6",
			arr:  []int{0, 3, 2, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArrayVer2(tt.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_validMountainArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validMountainArray([]int{0, 3, 2, 1})
	}
}

func Benchmark_validMountainArrayVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validMountainArrayVer2([]int{0, 3, 2, 1})
	}
}
