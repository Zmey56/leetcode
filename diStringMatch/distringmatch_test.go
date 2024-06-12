package diStringMatch

import "testing"

func Test_diStringMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{
			name: "Test case 1",
			s:    "IDID",
			want: []int{0, 4, 1, 3, 2},
		},
		{
			name: "Test case 2",
			s:    "III",
			want: []int{0, 1, 2, 3},
		},
		{
			name: "Test case 3",
			s:    "DDI",
			want: []int{3, 2, 0, 1},
		},
		{
			name: "Test case 4",
			s:    "I",
			want: []int{0, 1},
		},
		{
			name: "Test case 5",
			s:    "D",
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diStringMatch(tt.s); !equal(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_diStringMatchVer2(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{
			name: "Test case 1",
			s:    "IDID",
			want: []int{0, 4, 1, 3, 2},
		},
		{
			name: "Test case 2",
			s:    "III",
			want: []int{0, 1, 2, 3},
		},
		{
			name: "Test case 3",
			s:    "DDI",
			want: []int{3, 2, 0, 1},
		},
		{
			name: "Test case 4",
			s:    "I",
			want: []int{0, 1},
		},
		{
			name: "Test case 5",
			s:    "D",
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diStringMatchVer2(tt.s); !equal(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
