package groupAnagrams

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "Example 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name: "Empty input",
			strs: []string{},
			want: [][]string{},
		},
		{
			name: "Single element",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: "No anagrams",
			strs: []string{"a", "b", "c"},
			want: [][]string{
				{"a"},
				{"b"},
				{"c"},
			},
		},
		{
			name: "Duplicates",
			strs: []string{"abc", "bca", "abc"},
			want: [][]string{
				{"abc", "bca", "abc"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)

			gotNorm := normalize(got)
			wantNorm := normalize(tt.want)

			if !reflect.DeepEqual(gotNorm, wantNorm) {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func normalize(groups [][]string) map[string]int {
	result := make(map[string]int)

	for _, group := range groups {
		sorted := append([]string{}, group...)
		sort.Strings(sorted)
		key := strings.Join(sorted, "|")
		result[key]++
	}

	return result
}
