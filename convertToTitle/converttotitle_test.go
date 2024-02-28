package convertToTitle

import "testing"

//Test for function convertToTitle

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		columnNumber int
		want         string
	}{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},
		{2147483647, "FXSHRXW"},
	}

	for _, test := range tests {
		if got := convertToTitle(test.columnNumber); got != test.want {
			t.Errorf("convertToTitle(%v) = %v; want %v", test.columnNumber, got, test.want)
		}
	}
}
