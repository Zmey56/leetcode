package isPerfectSquare

import "testing"

// Test for function isPerfectSquare

func TestIsPerfectSquare(t *testing.T) {
	tests := []struct {
		num      int
		expected bool
	}{
		{16, true},
		{14, false},
		{1, true},
		{4, true},
		{9, true},
		{25, true},
		{36, true},
		{49, true},
		{64, true},
		{81, true},
		{100, true},
		{121, true},
		{144, true},
		{169, true},
		{196, true},
		{225, true},
	}

	for _, test := range tests {
		result := isPerfectSquare(test.num)
		if result != test.expected {
			t.Errorf("Test failed: expected %v, got %v", test.expected, result)
		}
	}
}
