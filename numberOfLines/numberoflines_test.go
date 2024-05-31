package numberOfLines

import "testing"

func Test_numberOfLines(t *testing.T) {
	tests := []struct {
		name   string
		widths []int
		s      string
		want   []int
	}{
		{
			name: "Test 1",
			widths: []int{
				10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
				10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
				10, 10, 10, 10, 10, 10,
			},
			s:    "abcdefghijklmnopqrstuvwxyz",
			want: []int{3, 60},
		},
		{
			name: "Test 2",
			widths: []int{
				4, 10, 10, 10, 10, 10, 10, 10, 10, 10,
				10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
				10, 10, 10, 10, 10, 10,
			},
			s:    "bbbcccdddaaa",
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfLines(tt.widths, tt.s); !equal(got, tt.want) {
				t.Errorf("numberOfLines() = %v, want %v", got, tt.want)
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
