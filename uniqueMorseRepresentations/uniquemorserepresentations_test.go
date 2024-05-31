package uniqueMorseRepresentations

import "testing"

func Test_uniqueMorseRepresentations(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{
			name:  "Example 1",
			words: []string{"gin", "zen", "gig", "msg"},
			want:  2,
		},
		{
			name:  "Test 1",
			words: []string{"a"},
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueMorseRepresentations(tt.words); got != tt.want {
				t.Errorf("uniqueMorseRepresentations() = %v, want %v", got, tt.want)
			}
		})
	}
}
