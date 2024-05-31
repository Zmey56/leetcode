package largeGroupPositions

import (
	"reflect"
	"testing"
)

func Test_largeGroupPositions(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]int
	}{
		{
			name: "Example 1",
			s:    "abbxxxxzzy",
			want: [][]int{{3, 6}},
		},
		{
			name: "Example 2",
			s:    "abc",
			want: [][]int{},
		},
		{
			name: "Example 3",
			s:    "abcdddeeeeaabbbcd",
			want: [][]int{{3, 5}, {6, 9}, {12, 14}},
		},
		{
			name: "Example 4",
			s:    "aaa",
			want: [][]int{{0, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largeGroupPositions(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_largeGroupPositions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largeGroupPositions("abbxxxxzzy")
	}
}

func Benchmark_largeGroupPositionsVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largeGroupPositionsVer2("abbxxxxzzy")
	}
}
