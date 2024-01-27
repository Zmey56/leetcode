package countAndSay

import "testing"

//Test for function countAndSay

func TestCounAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string // want is the expected result from calling countAndSay
	}{
		{
			name: "Test #1",
			args: args{
				n: 1,
			},
			want: "1",
		},
		{
			name: "Test #2",
			args: args{
				n: 2,
			},
			want: "11",
		},
		{
			name: "Test #3",
			args: args{
				n: 3,
			},
			want: "21",
		},
		{
			name: "Test #4",
			args: args{
				n: 4,
			},
			want: "1211",
		},
		{
			name: "Test #5",
			args: args{
				n: 5,
			},
			want: "111221",
		},
		{
			name: "Test #6",
			args: args{
				n: 6,
			},
			want: "312211",
		},
		{
			name: "Test #7",
			args: args{
				n: 7,
			},
			want: "13112221",
		},
		{
			name: "Test #8",
			args: args{
				n: 8,
			},
			want: "1113213211",
		},
		{
			name: "Test #9",
			args: args{
				n: 9,
			},
			want: "31131211131221",
		},
		{
			name: "Test #10",
			args: args{
				n: 10,
			},
			want: "13211311123113112211",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
