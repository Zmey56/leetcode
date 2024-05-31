package nextGreatestLetter

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	tests := []struct {
		name    string
		letters []byte
		target  byte
		want    byte
	}{
		{"Example 1", []byte{'c', 'f', 'j'}, 'a', 'c'},
		{"Example 2", []byte{'c', 'f', 'j'}, 'c', 'f'},
		{"Example 3", []byte{'c', 'f', 'j'}, 'd', 'f'},
		{"Example 4", []byte{'c', 'f', 'j'}, 'g', 'j'},
		{"Example 5", []byte{'c', 'f', 'j'}, 'j', 'c'},
		{"Example 6", []byte{'c', 'f', 'j'}, 'k', 'c'},
		{"Example 7", []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}, 'z', 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.letters, tt.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
