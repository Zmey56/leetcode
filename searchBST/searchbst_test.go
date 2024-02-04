package searchBST

import (
	"reflect"
	"testing"
)

//Test for function searchBST

func TestSearchBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{
			name: "Test 1",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 7},
			},
			val: 2,
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name: "Test 2",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 7},
			},
			val:  5,
			want: nil,
		},
	}
	for _, test := range tests {
		got := searchBST(test.root, test.val)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("searchBST(%v, %v) = %v; want %v", test.root, test.val, got, test.want)
		}
	}
}
