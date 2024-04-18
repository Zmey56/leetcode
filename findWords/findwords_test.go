package findWords

import "testing"

// Test for function findWords
func Test_findWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Test case 1",
			args{[]string{"Hello", "Alaska", "Dad", "Peace"}},
			[]string{"Alaska", "Dad"},
		},
		{
			"Test case 2",
			args{[]string{"omk"}},
			[]string{},
		},
		{
			"Test case 3",
			args{[]string{"adsdf", "sfd"}},
			[]string{"adsdf", "sfd"},
		},
		{
			"Test case 4",
			args{[]string{"a", "b"}},
			[]string{"a", "b"},
		},
		{
			"Test case 5",
			args{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}},
			[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.words); !compareStringSlices(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
