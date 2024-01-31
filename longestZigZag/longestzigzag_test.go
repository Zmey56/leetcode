package longestZigZag

import "testing"

//Test for function longestZigZag

func TestLongestZigZag(t *testing.T) {
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
					Val:   1,
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   1,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 1},
					},
					Right: &TreeNode{Val: 1},
				},
			},
			want: 3,
		},
		{
			name: "Test 2",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:  1,
						Left: &TreeNode{Val: 1},
					},
				},
			},
			want: 2,
		},
		{
			name: "Test 3",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  1,
						Left: &TreeNode{Val: 1},
					},
				},
			},
			want: 1,
		},
		{
			name: "Test 4",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:  1,
						Left: &TreeNode{Val: 1},
					},
				},
			},
			want: 3,
		},
	}

	for _, test := range tests {
		got := longestZigZag(test.root)
		if got != test.want {
			t.Errorf("longestZigZag(%v) = %d; want %d", test.name, got, test.want)
		}
	}
}
