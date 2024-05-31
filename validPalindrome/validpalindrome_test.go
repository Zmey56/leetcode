package validPalindrome

import "testing"

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				s: "aba",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				s: "abca",
			},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{
				s: "abc",
			},
			want: false,
		},
		{
			name: "Test case 4",
			args: args{
				s: "a",
			},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{
				s: "ab",
			},
			want: true,
		},
		{
			name: "Test case 6",
			args: args{
				s: "abba",
			},
			want: true,
		},
		{
			name: "Test case 7",
			args: args{
				s: "abcba",
			},
			want: true,
		},
		{
			name: "Test case 8",
			args: args{
				s: "abcdba",
			},
			want: true,
		},
		{
			name: "Test case 9",
			args: args{
				s: "abcdeba",
			},
			want: false,
		},
		{
			name: "Test case 10",
			args: args{
				s: "abccba",
			},
			want: true,
		},
		{
			name: "Test case 11",
			args: args{
				s: "abccdba",
			},
			want: true,
		},
		{
			name: "Test case 12",
			args: args{
				s: "abccdeba",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
