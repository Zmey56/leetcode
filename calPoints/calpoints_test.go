package calPoints

import "testing"

func Test_calPoints(t *testing.T) {
	type args struct {
		ops []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				ops: []string{"5", "2", "C", "D", "+"},
			},
			want: 30,
		},
		{
			name: "Test case 2",
			args: args{
				ops: []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			},
			want: 27,
		},
		{
			name: "Test case 3",
			args: args{
				ops: []string{"1"},
			},
			want: 1,
		},
		{
			name: "Test case 4",
			args: args{
				ops: []string{"1", "2"},
			},
			want: 3,
		},
		{
			name: "Test case 5",
			args: args{
				ops: []string{"1", "2", "3"},
			},
			want: 6,
		},
		{
			name: "Test case 6",
			args: args{
				ops: []string{"1", "2", "3", "C"},
			},
			want: 3,
		},
		{
			name: "Test case 7",
			args: args{
				ops: []string{"1", "2", "3", "C", "D"},
			},
			want: 7,
		},
		{
			name: "Test case 8",
			args: args{
				ops: []string{"1", "2", "3", "C", "D", "4"},
			},
			want: 11,
		},
		{
			name: "Test case 9",
			args: args{
				ops: []string{"1", "2", "3", "C", "D", "4", "5"},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calPoints(tt.args.ops); got != tt.want {
				t.Errorf("calPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
