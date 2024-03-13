package firstBadVersion

import "testing"

//Test for firstBadVersion

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Test Case 1",
			n:    5,
			want: 4,
		},
		{
			name: "Test Case 2",
			n:    1,
			want: 1,
		},
		{
			name: "Test Case 3",
			n:    10,
			want: 10,
		},
		{
			name: "Test Case 4",
			n:    100,
			want: 100,
		},
		{
			name: "Test Case 5",
			n:    1000,
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstBadVersion(tt.n)
			if got != tt.want {
				t.Errorf("Expected %v but got %v", tt.want, got)
			}
		})
	}
}
