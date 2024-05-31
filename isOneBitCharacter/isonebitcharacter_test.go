package isOneBitCharacter

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	tests := []struct {
		name string
		bits []int
		want bool
	}{
		{"Example 1", []int{1, 0, 0}, true},
		{"Example 2", []int{1, 1, 1, 0}, false},
		{"Example 3", []int{1, 1, 0}, true},
		{"Example 4", []int{0}, true},
		{"Example 5", []int{1, 0, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneBitCharacter(tt.bits); got != tt.want {
				t.Errorf("isOneBitCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
