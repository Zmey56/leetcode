package mostCommonWord

import "testing"

func Test_mostCommonWord(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		banned []string
		want   string
	}{
		{
			name:   "testcase 1",
			input:  "Bob hit a ball, the hit BALL flew far after it was hit.",
			banned: []string{"hit"},
			want:   "ball",
		},
		{
			name:   "testcase 2",
			input:  "a.",
			banned: []string{},
			want:   "a",
		},
		{
			name:   "testcase 3",
			input:  "Bob. hIt, baLl",
			banned: []string{"bob", "hit"},
			want:   "ball",
		},
		{
			name:   "testcase 4",
			input:  "a, a, a, a, b,b,b,c, c",
			banned: []string{"a"},
			want:   "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCommonWord(tt.input, tt.banned); got != tt.want {
				t.Errorf("mostCommonWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
