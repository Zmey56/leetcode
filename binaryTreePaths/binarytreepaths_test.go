package binaryTreePaths

import "testing"

// Test for binaryTreePaths
func TestBinaryTreePaths(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []string
	}{
		{
			name: "Test Case 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "Test Case 2",
			root: &TreeNode{
				Val: 1,
			},
			want: []string{"1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binaryTreePaths(tt.root)
			if len(got) != len(tt.want) {
				t.Errorf("Expected %v but got %v", tt.want, got)
			}
			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("Expected %v but got %v", tt.want, got)
				}
			}
		})
	}
}
