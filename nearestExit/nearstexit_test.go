package nearestExit

import "testing"

func Test_nearestExit(t *testing.T) {
	tests := []struct {
		name     string
		maze     [][]byte
		entrance []int
		want     int
	}{
		{
			name: "testcase 1",
			maze: [][]byte{
				{'+', '+', '.', '+'},
				{'.', '.', '.', '+'},
				{'+', '+', '+', '.'},
			},
			entrance: []int{1, 2},
			want:     1,
		},
		{
			name: "testcase 2",
			maze: [][]byte{
				{'.', '+'},
				{'+', '.'},
			},
			entrance: []int{1, 1},
			want:     -1,
		},
		{
			name: "testcase 3",
			maze: [][]byte{
				{'.', '.'},
			},
			entrance: []int{0, 0},
			want:     1,
		},
		{
			name: "testcase 4",
			maze: [][]byte{
				{'+', '+', '+'},
				{'.', '.', '.'},
				{'+', '+', '+'},
			},
			entrance: []int{1, 1},
			want:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestExit(tt.maze, tt.entrance); got != tt.want {
				t.Errorf("nearestExit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_nearestExit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nearestExit([][]byte{
			{'+', '+', '.', '+'},
			{'.', '.', '.', '+'},
			{'+', '+', '+', '.'},
		}, []int{1, 2})
	}
}

func Test_nearestExitVer2(t *testing.T) {
	tests := []struct {
		name     string
		maze     [][]byte
		entrance []int
		want     int
	}{
		{
			name: "testcase 1",
			maze: [][]byte{
				{'+', '+', '.', '+'},
				{'.', '.', '.', '+'},
				{'+', '+', '+', '.'},
			},
			entrance: []int{1, 2},
			want:     1,
		},
		{
			name: "testcase 2",
			maze: [][]byte{
				{'.', '+'},
				{'+', '.'},
			},
			entrance: []int{1, 1},
			want:     -1,
		},
		{
			name: "testcase 3",
			maze: [][]byte{
				{'.', '.'},
			},
			entrance: []int{0, 0},
			want:     1,
		},
		{
			name: "testcase 4",
			maze: [][]byte{
				{'+', '+', '+'},
				{'.', '.', '.'},
				{'+', '+', '+'},
			},
			entrance: []int{1, 1},
			want:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestExitVer2(tt.maze, tt.entrance); got != tt.want {
				t.Errorf("nearestExit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_nearestExitVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nearestExitVer2([][]byte{
			{'+', '+', '.', '+'},
			{'.', '.', '.', '+'},
			{'+', '+', '+', '.'},
		}, []int{1, 2})
	}
}
