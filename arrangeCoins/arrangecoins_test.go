package arrangeCoins

import "testing"

//Test for function arrangeCoins

func Test_arrangeCoins(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Test 1",
			n:    5,
			want: 2,
		},
		{
			name: "Test 2",
			n:    8,
			want: 3,
		},
		{
			name: "Test 3",
			n:    1,
			want: 1,
		},
		{
			name: "Test 4",
			n:    0,
			want: 0,
		},
		{
			name: "Test 5",
			n:    3,
			want: 2,
		},
		{
			name: "Test 6",
			n:    1804289383,
			want: 60070,
		},
	}
	for _, tt := range tests {
		if got := arrangeCoins(tt.n); got != tt.want {
			t.Errorf("Test - %v arrangeCoins() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
