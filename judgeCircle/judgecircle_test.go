package judgeCircle

import "testing"

// Test for function judgeCircle
func Test_judgeCircle(t *testing.T) {
	type args struct {
		moves string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//Test Case 1
		{
			name: "Test Case 1",
			args: args{
				moves: "UD",
			},
			want: true,
		},
		//Test Case 2
		{
			name: "Test Case 2",
			args: args{
				moves: "LL",
			},
			want: false,
		},
		//Test Case 3
		{
			name: "Test Case 3",
			args: args{
				moves: "UDLR",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeCircle(tt.args.moves); got != tt.want {
				t.Errorf("judgeCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
