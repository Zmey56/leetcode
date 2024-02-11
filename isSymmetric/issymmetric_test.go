package isSymmetric

import "testing"

//Test for function isSymmetric

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name  string
		input *TreeNode
		want  bool
	}{
		{
			name:  "Example 1",
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}},
			want:  true,
		},
		{
			name:  "Example 2",
			input: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			want:  false,
		},
		{
			name:  "Example 3",
			input: &TreeNode{},
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSymmetric(tt.input)
			if got != tt.want {
				t.Errorf("%v isSymmetric() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
