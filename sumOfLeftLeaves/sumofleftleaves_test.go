package sumOfLeftLeaves

import "testing"

//Test for function SumOfLeftLeaves

func Test_SumOfLeftLeaves(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: 24,
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			want: 0,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		if got := SumOfLeftLeaves(tt.root); got != tt.want {
			t.Errorf("SumOfLeftLeaves() = %v, want %v", got, tt.want)
		}
	}
}
