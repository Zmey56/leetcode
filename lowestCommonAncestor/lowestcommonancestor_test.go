package lowestCommonAncestor

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	test := []struct {
		name string
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			name: "Test 1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 6},
					Right: &TreeNode{Val: 2,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 8},
				},
			},
			p:    &TreeNode{Val: 5},
			q:    &TreeNode{Val: 1},
			want: &TreeNode{Val: 3},
		},
		{
			name: "Test 2",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 6},
					Right: &TreeNode{Val: 2,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 8},
				},
			},
			p:    &TreeNode{Val: 5},
			q:    &TreeNode{Val: 4},
			want: &TreeNode{Val: 5},
		},
	}
	for _, test := range test {
		got := lowestCommonAncestor(test.root, test.p, test.q)
		if got != test.want {
			t.Errorf("lowestCommonAncestor(%v, %v, %v) = %v; want %v", test.root, test.p, test.q, got, test.want)
		}
	}
}
