package hammingDistance

import "testing"

// Test for hammingDistance

func Test_hammingDistance(t *testing.T) {
	type args struct {
		x int
		y int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				x: 1,
				y: 4,
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				x: 3,
				y: 1,
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				x: 3,
				y: 3,
			},
			want: 0,
		},
		{
			name: "Test 4",
			args: args{
				x: 0,
				y: 0,
			},
			want: 0,
		},
		{
			name: "Test 5",
			args: args{
				x: 0,
				y: 1,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("hammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
