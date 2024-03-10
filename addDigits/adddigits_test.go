package addDigits

import "testing"

//Test for addDigits

func TestAddDigits(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{38, 2},
		{0, 0},
		{9, 9},
		{1, 1},
		{10, 1},
		{150, 1},
		{174, 1},
		{645, 1},
	}
	for _, test := range tests {
		if got := addDigits(test.num); got != test.want {
			t.Errorf("addDigits(%v) = %v; want %v", test.num, got, test.want)
		}
	}
}
