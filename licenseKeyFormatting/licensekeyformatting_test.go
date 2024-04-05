package licenseKeyFormatting

import "testing"

//Tetst for function licenseKeyFormatting

func Test_licenseKeyFormatting(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test case 1",
			args{"5F3Z-2e-9-w", 4},
			"5F3Z-2E9W",
		},
		{
			"Test case 2",
			args{"2-5g-3-J", 2},
			"2-5G-3J",
		},
		{
			"Test case 3",
			args{"2-5g-3-J", 3},
			"25-G3J",
		},
		{
			"Test case 4",
			args{"2-5g-3-J", 1},
			"2-5-G-3-J",
		},
		{
			"Test case 5",
			args{"2-5g-3-J", 4},
			"2-5G3J",
		},
		{
			"Test case 6",
			args{"2-5g-3-J", 5},
			"25G3J",
		},
		{
			"Test case 7",
			args{"2-5g-3-J", 6},
			"25G3J",
		},
		{
			"Test case 8",
			args{"2-5g-3-J", 7},
			"25G3J",
		},
		{
			"Test case 9",
			args{"2-5g-3-J", 8},
			"25G3J",
		},
		{
			"Test case 10",
			args{"2-5g-3-J", 9},
			"25G3J",
		},
		{
			"Test case 11",
			args{"2-5g-3-J", 10},
			"25G3J",
		},
		{
			"Test case 12",
			args{"2-5g-3-J", 11},
			"25G3J",
		},
		{
			"Test case 13",
			args{"-----", 3},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}
