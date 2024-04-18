package findMode

import "testing"

//Test for function findMode

func Test_findMode(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test case 1",
			args{&TreeNode{1, nil, nil}},
			[]int{1},
		},
		{
			"Test case 2",
			args{&TreeNode{1, &TreeNode{2, nil, nil}, nil}},
			[]int{1, 2},
		},
		{
			"Test case 3",
			args{&TreeNode{1, &TreeNode{1, nil, nil}, nil}},
			[]int{1},
		},
		{
			"Test case 4",
			args{&TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}},
			[]int{1},
		},
		{
			"Test case 5",
			args{&TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !compareIntSlices(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true

}
