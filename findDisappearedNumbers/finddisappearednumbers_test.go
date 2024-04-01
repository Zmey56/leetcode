package findDisappearedNumbers

import (
	"reflect"
	"testing"
)

//Test for function findDisappearedNumbers

func Test_findDisappearedNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Test 1",
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
		{
			name: "Test 2",
			nums: []int{1, 1},
			want: []int{2},
		},
		{
			name: "Test 3",
			nums: []int{1},
			want: []int{},
		},
		{
			name: "Test 4",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: []int{},
		},
		{
			name: "Test 5",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int{},
		},
		{
			name: "Test 6",
			nums: []int{1, 1, 1, 1, 1, 1, 1, 1},
			want: []int{2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		if got := findDisappearedNumbers(tt.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Test - %v findDisappearedNumbers() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
