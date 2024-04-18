package checkPerfectNumber

import "testing"

// Test for function checkPerfectNumber
func Test_checkPerfectNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test case 1",
			args{28},
			true,
		},
		{
			"Test case 2",
			args{6},
			true,
		},
		{
			"Test case 3",
			args{496},
			true,
		},
		{
			"Test case 4",
			args{8128},
			true,
		},
		{
			"Test case 5",
			args{2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPerfectNumber(tt.args.num); got != tt.want {
				t.Errorf("checkPerfectNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
