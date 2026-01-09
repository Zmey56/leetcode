package evalRPN

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "Test case 1",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			name:   "Test case 2",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			name:   "Test case 3",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
		{
			name:   "Test case 4",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := evalRPN(test.tokens)
			if got != test.want {
				t.Errorf("evalRPN(%v) = %v, want %v", test.tokens, got, test.want)
			}
		})
	}
}
