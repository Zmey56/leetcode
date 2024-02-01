package rightSideView

import "testing"

//Test for function rightSideView

func TestRightSideView(t *testing.T) {
	test := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "Test 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  5,
						Left: &TreeNode{Val: 7},
					},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "Test 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
			want: []int{1, 3, 4},
		},
		{
			name: "Test 3",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 5},
				},
			},
			want: []int{1, 3, 5},
		},
	}

	for _, test := range test {
		got := rightSideView(test.root)
		if !IsSameSlice(got, test.want) {
			t.Errorf("rightSideView(%v) = %v; want %v", test.root, got, test.want)
		}
	}
}

// Function IsSameSlice
func IsSameSlice(a, b []int) bool {
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
