package backspaceCompare

import "testing"

func Test_backspaceCompare(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "testcase 1",
			s:    "ab#c",
			t:    "ad#c",
			want: true,
		},
		{
			name: "testcase 2",
			s:    "ab##",
			t:    "c#d#",
			want: true,
		},
		{
			name: "testcase 3",
			s:    "a##c",
			t:    "#a#c",
			want: true,
		},
		{
			name: "testcase 4",
			s:    "a#c",
			t:    "b",
			want: false,
		},
		{
			name: "testcase 5",
			s:    "a#c",
			t:    "b",
			want: false,
		},
		{
			name: "testcase 6",
			s:    "a#c",
			t:    "b",
			want: false,
		},
		{
			name: "testcase 7",
			s:    "a#c",
			t:    "b",
			want: false,
		},
		{
			name: "testcase 8",
			s:    "a#c",
			t:    "b",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompareVer2(tt.s, tt.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Banchmark_backspaceCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		backspaceCompare("ab#c", "ad#c")
	}
}

func Banchmark_backspaceCompareVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		backspaceCompareVer2("ab#c", "ad#c")
	}
}
