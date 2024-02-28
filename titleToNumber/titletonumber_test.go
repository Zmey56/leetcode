package titleToNumber

import "testing"

// Test for function titleToNumber

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		columnTitle string
		want        int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
		{"FXSHRXW", 2147483647},
	}

	for _, test := range tests {
		if got := titleToNumber(test.columnTitle); got != test.want {
			t.Errorf("titleToNumber(%v) = %v; want %v", test.columnTitle, got, test.want)
		}
	}
}
