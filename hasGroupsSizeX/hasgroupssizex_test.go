package hasGroupsSizeX

import "testing"

func Test_hasGroupsSizeX(t *testing.T) {
	tests := []struct {
		name string
		deck []int
		want bool
	}{
		{
			name: "Test case 1",
			deck: []int{1, 2, 3, 4, 4, 3, 2, 1},
			want: true,
		},
		{
			name: "Test case 2",
			deck: []int{1, 1, 1, 2, 2, 2, 3, 3},
			want: false,
		},
		{
			name: "Test case 3",
			deck: []int{1},
			want: false,
		},
		{
			name: "Test case 4",
			deck: []int{1, 1},
			want: true,
		},
		{
			name: "Test case 5",
			deck: []int{1, 1, 2, 2, 2, 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasGroupsSizeX(tt.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX() = %v, want %v", got, tt.want)
			}
		})
	}
}
