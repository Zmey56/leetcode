package maxLevelSum

import "testing"

// Test for function maxLevelSum

func TestMaxLevelSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "Test 1",

			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: -8},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			want: 2,
		},
		{
			name: "Test 2",
			root: &TreeNode{
				Val: 989,
				Left: &TreeNode{
					Val:   10250,
					Left:  &TreeNode{Val: 98693},
					Right: &TreeNode{Val: -89388},
				},
				Right: &TreeNode{
					Val: -32127,
				},
			},
			want: 2,
		},
	}
	for _, test := range tests {
		got := maxLevelSum(test.root)
		if got != test.want {
			t.Errorf("maxLevelSum(%v) = %d; want %d", test.name, got, test.want)
		}
	}
}
