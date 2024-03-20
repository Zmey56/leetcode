package toHex

import "testing"

// Test for function toHex
func Test_toHex(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{
			num:  26,
			want: "1a",
		},
		{
			num:  -1,
			want: "ffffffff",
		},
		{
			num:  0,
			want: "0",
		},
	}
	for _, tt := range tests {
		if got := toHex(tt.num); got != tt.want {
			t.Errorf("toHex() = %v, want %v", got, tt.want)
		}
	}
}
