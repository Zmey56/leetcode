package flipAndInvertImage

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	tests := []struct {
		name  string
		image [][]int
		want  [][]int
	}{
		{
			name:  "Example 1",
			image: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			want:  [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			name:  "Example 2",
			image: [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			want:  [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipAndInvertImage(tt.image); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
