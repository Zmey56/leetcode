package repeatedSubstringPattern

import "testing"

//Test function for repeatedSubstringPattern

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				s: "abab",
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				s: "aba",
			},
			want: false,
		},
		{
			name: "Test 3",
			args: args{
				s: "abcabcabcabc",
			},
			want: true,
		},
		{
			name: "Test 4",
			args: args{
				s: "abcabcabc",
			},
			want: true,
		},
		{
			name: "Test 5",
			args: args{
				s: "a",
			},
			want: false,
		},
		{
			name: "Test 6",
			args: args{
				s: "ababba",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		if got := repeatedSubstringPattern(tt.args.s); got != tt.want {
			t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
		}
	}
}
