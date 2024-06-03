package uncommonFromSentences

import (
	"sort"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want []string
	}{
		{
			name: "Example 1",
			s1:   "this apple is sweet",
			s2:   "this apple is sour",
			want: []string{"sweet", "sour"},
		},
		{
			name: "Example 2",
			s1:   "apple apple",
			s2:   "banana",
			want: []string{"banana"},
		},
		{
			name: "Example 3",
			s1:   "s z z z s",
			s2:   "s z ejt",
			want: []string{"ejt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uncommonFromSentences(tt.s1, tt.s2); !deepEqualIgnoreOrder(got, tt.want) {
				t.Errorf("uncommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func deepEqualIgnoreOrder(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Benchmark_uncommonFromSentences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uncommonFromSentences("this apple is sweet", "this apple is sour")
	}
}

func Benchmark_uncommonFromSentences2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uncommonFromSentencesVer2("this apple is sweet", "this apple is sour")
	}
}
