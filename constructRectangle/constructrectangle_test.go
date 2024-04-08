package constructRectangle

import (
	"reflect"
	"testing"
)

//Test for function constructRectangle

func Test_constructRectangle(t *testing.T) {
	type args struct {
		area int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test case 1",
			args{4},
			[]int{2, 2},
		},
		{
			"Test case 2",
			args{37},
			[]int{37, 1},
		},
		{
			"Test case 3",
			args{122122},
			[]int{427, 286},
		},
		{
			"Test case 4",
			args{1},
			[]int{1, 1},
		},
		{
			"Test case 5",
			args{10000000},
			[]int{3200, 3125},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructRectangle(tt.args.area); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
