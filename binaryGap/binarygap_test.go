package binaryGap

import "testing"

func Test_binaryGap(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    22,
			want: 2,
		},
		{
			name: "Example 2",
			n:    5,
			want: 2,
		},
		{
			name: "Example 3",
			n:    6,
			want: 1,
		},
		{
			name: "Example 4",
			n:    8,
			want: 0,
		},
		{
			name: "Example 5",
			n:    1,
			want: 0,
		},
		{
			name: "Example 6",
			n:    13,
			want: 2,
		},
		{
			name: "Example 7",
			n:    123456,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.n); got != tt.want {
				t.Errorf("binaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
