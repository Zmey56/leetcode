package isUgly

import "testing"

// Test for isUgly
func TestIsUgly(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{6, true},
		{8, true},
		{14, false},
		{1, true},
		{0, false},
	}
	for _, test := range tests {
		if got := isUgly(test.num); got != test.want {
			t.Errorf("isUgly(%v) = %v; want %v", test.num, got, test.want)
		}
	}
}
