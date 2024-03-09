package isPowerOfTwo

import "testing"

//Tets for isPowerOfTwo

func TesrIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{
			name: "Test Case 1",
			num:  1,
			want: true,
		},
		{
			name: "Test Case 2",
			num:  16,
			want: true,
		},
		{
			name: "Test Case 3",
			num:  3,
			want: false,
		},
		{
			name: "Test Case 4",
			num:  4,
			want: true,
		},
		{
			name: "Test Case 5",
			num:  5,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPowerOfTwo(tt.num)
			if got != tt.want {
				t.Errorf("Expected %v but got %v", tt.want, got)
			}
		})
	}
}
