package canVisitAllRooms

import "testing"

// test for function canVisitAllRooms
func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  bool
	}{
		{
			name: "Test 1",
			rooms: [][]int{
				{1},
				{2},
				{3},
				{},
			},
			want: true,
		},
		{
			name: "Test 2",
			rooms: [][]int{
				{1, 3},
				{3, 0, 1},
				{2},
				{0},
			},
			want: false,
		},
	}
	for _, test := range tests {
		got := canVisitAllRooms(test.rooms)
		if got != test.want {
			panic("Test failed")
		}
	}

}
