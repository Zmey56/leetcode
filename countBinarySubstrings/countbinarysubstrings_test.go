package countBinarySubstrings

import "testing"

func Test_countBinarySubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"00110011"}, 6},
		{"2", args{"10101"}, 4},
		{"3", args{"00110"}, 3},
		{"4", args{"00100"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBinarySubstrings(tt.args.s); got != tt.want {
				t.Errorf("countBinarySubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
