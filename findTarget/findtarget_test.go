package findTarget

import "testing"

// Test for function findTarget
func Test_findTarget(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//Test Case 1
		{
			name: "Test Case 1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 9,
			},
			want: true,
		},
		//Test Case 2
		{
			name: "Test Case 2",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				k: 28,
			},
			want: false,
		},
		//Test Case 3
		{
			name: "Test Case 3",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				k: 4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
