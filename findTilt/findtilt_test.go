package findTilt

import "testing"

//Test for function findTilt

func Test_findTilt(t *testing.T) {
	// Test case 1
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	if got := findTilt(root1); got != 1 {
		t.Errorf("findTilt() = %v, want %v", got, 1)
	}

	// Test case 2
	root2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	if got := findTilt(root2); got != 15 {
		t.Errorf("findTilt() = %v, want %v", got, 15)
	}
}
