package shortestToChar

import (
	"reflect"
	"testing"
)

func Test_shortestToChar(t *testing.T) {
	tests := []struct {
		name  string
		input string
		c     byte
		want  []int
	}{
		{
			name:  "testcase 1",
			input: "loveleetcode",
			c:     'e',
			want:  []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			name:  "testcase 2",
			input: "aaab",
			c:     'b',
			want:  []int{3, 2, 1, 0},
		},
		{
			name:  "testcase 3",
			input: "aaba",
			c:     'b',
			want:  []int{2, 1, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestToChar(tt.input, tt.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_shortestToChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shortestToChar("loveleetcode", 'e')
	}
}
