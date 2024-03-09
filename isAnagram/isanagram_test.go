package isAnagram

import "testing"

//Test for isAnagram

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Test Case 1",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "Test Case 2",
			s:    "rat",
			t:    "car",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("Expected %v but got %v", tt.want, got)
			}
		})
	}
}
