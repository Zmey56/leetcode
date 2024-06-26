package reverseWordsTwo

import "testing"

// Test for function reverseWords

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test case 1", args{"Let's take LeetCode contest"}, "s'teL ekat edoCteeL tsetnoc"},
		{"Test case 2", args{"God Ding"}, "doG gniD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
