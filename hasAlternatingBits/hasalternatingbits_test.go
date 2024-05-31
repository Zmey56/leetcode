package hasAlternatingBits

import "testing"

func Test_hasAlternatingBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{5}, true},
		{"2", args{7}, false},
		{"3", args{11}, false},
		{"4", args{10}, true},
		{"5", args{4}, false},
		{"6", args{3}, false},
		{"7", args{1}, true},
		{"8", args{2}, true},
		{"9", args{6}, false},
		{"10", args{8}, false},
		{"11", args{9}, false},
		{"12", args{12}, false},
		{"13", args{13}, false},
		{"14", args{14}, false},
		{"15", args{15}, false},
		{"16", args{16}, false},
		{"17", args{17}, false},
		{"18", args{18}, false},
		{"19", args{19}, false},
		{"20", args{20}, false},
		{"21", args{21}, true},
		{"22", args{22}, false},
		{"23", args{23}, false},
		{"24", args{24}, false},
		{"25", args{25}, false},
		{"26", args{26}, false},
		{"27", args{27}, false},
		{"28", args{28}, false},
		{"29", args{29}, false},
		{"30", args{30}, false},
		{"31", args{31}, false},
		{"32", args{32}, false},
		{"33", args{33}, false},
		{"34", args{34}, false},
		{"35", args{35}, false},
		{"36", args{36}, false},
		{"37", args{37}, false},
		{"38", args{38}, false},
		{"39", args{39}, false},
		{"40", args{40}, false},
		{"41", args{41}, false},
		{"42", args{42}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAlternatingBits(tt.args.n); got != tt.want {
				t.Errorf("hasAlternatingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
