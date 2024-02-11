package inorderTraversal

import "testing"

// Test for function inorderTraversal

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name  string
		input *TreeNode
		want  []int
	}{
		{
			name:  "Example 1",
			input: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			want:  []int{1, 3, 2},
		},
		{
			name:  "Example 2",
			input: &TreeNode{},
			want:  []int{},
		},
		{
			name:  "Example 3",
			input: &TreeNode{Val: 1},
			want:  []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inorderTraversal(tt.input)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("%v inorderTraversal() = %v, want %v", tt.name, got, tt.want)
				}
			}
		})
	}
}
