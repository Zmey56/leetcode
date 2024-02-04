package findCircleNum

import "testing"

// Test for function findCircleNum

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  int
	}{
		{
			name: "Test 1",
			rooms: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			want: 2,
		},
		{
			name: "Test 2",
			rooms: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: 3,
		},
	}

	for _, test := range tests {
		got := findCircleNum(test.rooms)
		if got != test.want {
			t.Errorf("findCircleNum(%v) = %d; want %d", test.rooms, got, test.want)
		}
	}
}
