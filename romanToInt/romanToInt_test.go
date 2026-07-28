package main

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{"simple addition", "III", 3},
		{"simple subtraction", "IV", 4},
		{"subtraction with nine", "IX", 9},
		{"complex calculation", "LVIII", 58}, // L = 50, V = 5, III = 3
		{"year 1994", "MCMXCIV", 1994},       // M=1000, CM=900, XC=90, IV=4
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.s)

			if got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
