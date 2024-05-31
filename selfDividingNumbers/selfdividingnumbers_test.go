package selfDividingNumbers

import (
	"reflect"
	"testing"
)

func Test_selfDividingNumbers(t *testing.T) {
	tests := []struct {
		name  string
		left  int
		right int
		want  []int
	}{
		{"Example 1", 1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{"Example 2", 47, 85, []int{48, 55, 66, 77}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selfDividingNumbers(tt.left, tt.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selfDividingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
