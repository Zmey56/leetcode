package minWindow

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		expect string
	}{
		{
			s:      "ADOBECODEBANC",
			t:      "ABC",
			expect: "BANC",
		},
		{
			s:      "a",
			t:      "a",
			expect: "a",
		},
		{
			s:      "a",
			t:      "aa",
			expect: "",
		},
		{
			s:      "aaflslflsldkalskaaa",
			t:      "aaa",
			expect: "aaa",
		},
	}

	for _, tt := range tests {
		result := minWindow(tt.s, tt.t)
		if result != tt.expect {
			t.Errorf("For input s='%s' and t='%s', expected '%s' but got '%s'", tt.s, tt.t, tt.expect, result)
		}
	}
}
