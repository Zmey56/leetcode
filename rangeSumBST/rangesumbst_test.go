package rangeSumBST

import "testing"

func Test_rangeSumBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		low  int
		high int
		want int
	}{
		{
			name: "Test case 1",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 15,
					Right: &TreeNode{
						Val: 18,
					},
				},
			},
			low:  7,
			high: 15,
			want: 32,
		},
		{
			name: "Test case 2",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 18,
					},
				},
			},
			low:  6,
			high: 10,
			want: 23,
		},
		{
			name: "Test case 3",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 18,
					},
				},
			},
			low:  1,
			high: 20,
			want: 85,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBST(tt.root, tt.low, tt.high); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
