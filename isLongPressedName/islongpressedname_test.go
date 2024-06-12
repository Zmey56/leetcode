package isLongPressedName

import "testing"

func Test_isLongPressedName(t *testing.T) {
	tests := []struct {
		name  string
		name1 string
		name2 string
		want  bool
	}{
		{
			name:  "Test case 1",
			name1: "alex",
			name2: "aaleex",
			want:  true,
		},
		{
			name:  "Test case 2",
			name1: "saeed",
			name2: "ssaaedd",
			want:  false,
		},
		{
			name:  "Test case 3",
			name1: "leelee",
			name2: "lleeelee",
			want:  true,
		},
		{
			name:  "Test case 4",
			name1: "laiden",
			name2: "laiden",
			want:  true,
		},
		{
			name:  "Test case 5",
			name1: "pyplrz",
			name2: "ppyypllr",
			want:  false,
		},
		{
			name:  "Test case 6",
			name1: "rick",
			name2: "kric",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.name1, tt.name2); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
