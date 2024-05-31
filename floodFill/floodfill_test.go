package floodFill

import (
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	tests := []struct {
		name   string
		image  [][]int
		sr     int
		sc     int
		color  int
		output [][]int
	}{
		{"Example 1", [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{"Example 2", [][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1, [][]int{{0, 0, 0}, {0, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.image, tt.sr, tt.sc, tt.color); !reflect.DeepEqual(got, tt.output) {
				t.Errorf("floodFill() = %v, want %v", got, tt.output)
			}
		})
	}
}
