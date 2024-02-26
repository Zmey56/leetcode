package maxProfit

import "testing"

// Test for function maxProfit
func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{"Testcase 1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"Testcase 2", []int{1, 2, 3, 4, 5}, 4},
		{"Testcase 3", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxProfit(test.prices)
			if got != test.expected {
				t.Errorf("maxProfit(%v) = %v, want %v", test.prices, got, test.expected)
			}
		})
	}
}
