package checkRecord

import "testing"

// Test for function checkRecord

func Test_checkRecord(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		result bool
	}{
		{"Test 1", "PPALLP", true},
		{"Test 2", "PPALLL", false},
		{"Test 3", "LALL", true},
		{"Test 4", "ALLAPPL", false},
	}

	for _, tt := range tests {
		if got := checkRecord(tt.s); got != tt.result {
			t.Errorf("%v checkRecord(%v) = %v, want %v", tt.name, tt.s, got, tt.result)
		}
	}
}
