package convertToBase7

import "testing"

// Test for function convertToBase7

func Test_convertToBase7(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test case 1",
			args{100},
			"202",
		},
		{
			"Test case 2",
			args{-7},
			"-10",
		},
		{
			"Test case 3",
			args{0},
			"0",
		},
		{
			"Test case 4",
			args{7},
			"10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBase7(tt.args.num); got != tt.want {
				t.Errorf("convertToBase7() = %v, want %v", got, tt.want)
			}
		})
	}
}
