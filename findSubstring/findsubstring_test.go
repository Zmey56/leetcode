package findSubstring

import "testing"

func TestFindSubstring(t *testing.T) {
	test := []struct {
		s     string
		words []string
		want  []int
	}{
		{
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
			want:  []int{0, 9},
		},
		{
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "word"},
			want:  []int{},
		},
		{
			s:     "barfoofoobarthefoobarman",
			words: []string{"bar", "foo", "the"},
			want:  []int{6, 9, 12},
		},
		{
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "good"},
			want:  []int{8},
		},
	}

	for _, test := range test {
		got := findSubstring(test.s, test.words)
		if !equal(got, test.want) {
			t.Errorf("findSubstring(%s, %v) = %v, want %v", test.s, test.words, got, test.want)
		}
	}
}

// check two slices are equal or not of any order
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
