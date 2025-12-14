package findSubstringV2

import "testing"

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		s      string
		words  []string
		expect []int
	}{
		{
			s:      "barfoothefoobarman",
			words:  []string{"foo", "bar"},
			expect: []int{0, 9},
		},
		{
			s:      "wordgoodgoodgoodbestword",
			words:  []string{"word", "good", "best", "word"},
			expect: []int{},
		},
		{
			s:      "barfoofoobarthefoobarman",
			words:  []string{"bar", "foo", "the"},
			expect: []int{6, 9, 12},
		},
	}

	for _, tt := range tests {
		result := findSubstring(tt.s, tt.words)
		if len(result) != len(tt.expect) {
			t.Errorf("For input s='%s' and words=%v, expected %v but got %v", tt.s, tt.words, tt.expect, result)
			continue
		}
		for i := range result {
			if result[i] != tt.expect[i] {
				t.Errorf("For input s='%s' and words=%v, expected %v but got %v", tt.s, tt.words, tt.expect, result)
				break
			}
		}
	}
}
