package isSameTree

import "testing"

//Test for function isSameTree

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    []int
		q    []int
		want bool
	}{
		{
			name: "Example 1",
			p:    []int{1, 2, 3},
			q:    []int{1, 2, 3},
			want: true,
		},
		{
			name: "Example 2",
			p:    []int{1, 2},
			q:    []int{1, -1, 2},
			want: false,
		},
		{
			name: "Example 3",
			p:    []int{1, 2, 1},
			q:    []int{1, 1, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSameTree(createTreeNode(tt.p), createTreeNode(tt.q))
			if got != tt.want {
				t.Errorf("%v isSameTree() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

// convert from slice string to *TreeNode
func createTreeNode(input []int) *TreeNode {
	if len(input) == 0 {
		return nil
	}
	root := &TreeNode{Val: input[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(input); i++ {
		node := queue[0]
		queue = queue[1:]
		if input[i] != -1 {
			node.Left = &TreeNode{Val: input[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(input) && input[i] != -1 {
			node.Right = &TreeNode{Val: input[i]}
			queue = append(queue, node.Right)
		}
	}
	return root
}
