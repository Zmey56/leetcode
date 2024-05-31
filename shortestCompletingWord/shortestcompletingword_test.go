package shortestCompletingWord

import "testing"

func Test_shortestCompletingWord(t *testing.T) {
	tests := []struct {
		name     string
		license  string
		words    []string
		expected string
	}{
		{"Example 1", "1s3 PSt", []string{"step", "steps", "stripe", "stepple"}, "steps"},
		{"Example 2", "1s3 456", []string{"looks", "pest", "stew", "show"}, "pest"},
		{"Example 3", "Ah71752", []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"}, "husband"},
		{"Example 4", "OgEu755", []string{"enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"}, "enough"},
		{"Example 5", "iMSlpe4", []string{"claim", "consumer", "student", "camera", "public", "never", "wonder", "simple", "thought", "use"}, "simple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestCompletingWord(tt.license, tt.words); got != tt.expected {
				t.Errorf("shortestCompletingWord() = %v, want %v", got, tt.expected)
			}
		})
	}
}
