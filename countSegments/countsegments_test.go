package countSegments

import "testing"

// Test for function countSegments

func Test_countSegments(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Test 1",
			s:    "Hello, my name is John",
			want: 5,
		},
		{
			name: "Test 2",
			s:    "Hello",
			want: 1,
		},
		{
			name: "Test 3",
			s:    "love live! mu'sic forever",
			want: 4,
		},
		{
			name: "Test 4",
			s:    "",
			want: 0,
		},
		{
			name: "Test 5",
			s:    "     ",
			want: 0,
		},
		{
			name: "Test 6",
			s:    ", , , ,        a, eaefa",
			want: 6,
		},
	}
	for _, tt := range tests {
		if got := countSegments(tt.s); got != tt.want {
			t.Errorf("Test - %v countSegments() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
