package findComplement

import "testing"

//Test for function findComplement

func Test_findComplement(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test case 1",
			args{5},
			2,
		},
		{
			"Test case 2",
			args{1},
			0,
		},
		{
			"Test case 3",
			args{2},
			1,
		},
		{
			"Test case 4",
			args{3},
			0,
		},
		{
			"Test case 5",
			args{4},
			3,
		},
		{
			"Test case 6",
			args{6},
			1,
		},
		{
			"Test case 7",
			args{7},
			0,
		},
		{
			"Test case 8",
			args{8},
			7,
		},
		{
			"Test case 9",
			args{9},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findComplement(tt.args.num); got != tt.want {
				t.Errorf("findComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
