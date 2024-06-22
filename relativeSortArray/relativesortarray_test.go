package relativeSortArray

import (
	"reflect"
	"testing"
)

func Test_relativeSortArray(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			name: "Test 1",
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			name: "Test 2",
			arr1: []int{28, 6, 22, 8, 44, 17},
			arr2: []int{22, 28, 8, 6},
			want: []int{22, 28, 8, 6, 17, 44},
		},
		{
			name: "Test 3",
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relativeSortArray(tt.arr1, tt.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relativeSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_relativeSortArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})
	}
}

func Test_relativeSortArrayVer2(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			name: "Test 1",
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			name: "Test 2",
			arr1: []int{28, 6, 22, 8, 44, 17},
			arr2: []int{22, 28, 8, 6},
			want: []int{22, 28, 8, 6, 17, 44},
		},
		{
			name: "Test 3",
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relativeSortArrayVer2(tt.arr1, tt.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relativeSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_relativeSortArrayVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		relativeSortArrayVer2([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})
	}
}

func Test_relativeSortArrayVer3(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			name: "Test 1",
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			name: "Test 2",
			arr1: []int{28, 6, 22, 8, 44, 17},
			arr2: []int{22, 28, 8, 6},
			want: []int{22, 28, 8, 6, 17, 44},
		},
		{
			name: "Test 3",
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relativeSortArrayVer3(tt.arr1, tt.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relativeSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_relativeSortArrayVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		relativeSortArrayVer3([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})
	}
}
