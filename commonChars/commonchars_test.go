package commonChars

import "testing"

func Test_commonChars(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{
			name:  "Test case 1",
			words: []string{"bella", "label", "roller"},
			want:  []string{"e", "l", "l"},
		},
		{
			name:  "Test case 2",
			words: []string{"cool", "lock", "cook"},
			want:  []string{"c", "o"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.words); !deepEqualIgnoreOrder(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func deepEqualIgnoreOrder(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, v := range a {
		if !contains(b, v) {
			return false
		}
	}

	for _, v := range b {
		if !contains(a, v) {
			return false
		}
	}

	return true
}

func contains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}
