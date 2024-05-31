package imageSmoother

import (
	"reflect"
	"testing"
)

func Test_imageSmoother(t *testing.T) {
	type args struct {
		img [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//Test Case 1
		{
			name: "Test Case 1",
			args: args{
				img: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		//Test Case 2
		{
			name: "Test Case 2",
			args: args{
				img: [][]int{
					{1},
				},
			},
			want: [][]int{
				{1},
			},
		},
		//Test Case 3
		{
			name: "Test Case 3",
			args: args{
				img: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: [][]int{
				{3, 3, 4},
				{4, 5, 5},
				{6, 6, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imageSmoother(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageSmoother() = %v, want %v", got, tt.want)
			}
		})
	}
}
