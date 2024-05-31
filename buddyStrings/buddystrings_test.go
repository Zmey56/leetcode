package buddyStrings

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				s:    "ab",
				goal: "ba",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s:    "ab",
				goal: "ab",
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				s:    "aa",
				goal: "aa",
			},
			want: true,
		},
		{
			name: "Example 4",
			args: args{
				s:    "aaaaaaabc",
				goal: "aaaaaaacb",
			},
			want: true,
		},
		{
			name: "Example 5",
			args: args{
				s:    "",
				goal: "aa",
			},
			want: false,
		},
		{
			name: "Example 6",
			args: args{
				s:    "ab",
				goal: "ac",
			},
			want: false,
		},
		{
			name: "Example 7",
			args: args{
				s:    "ab",
				goal: "ca",
			},
			want: false,
		},
		{
			name: "Example 8",
			args: args{
				s:    "ab",
				goal: "cd",
			},
			want: false,
		},
		{
			name: "Example 9",
			args: args{
				s:    "ab",
				goal: "ba",
			},
			want: true,
		},
		{
			name: "Example 10",
			args: args{
				s:    "ab",
				goal: "ab",
			},
			want: false,
		},
		{
			name: "Example 11",
			args: args{
				s:    "aa",
				goal: "aa",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
