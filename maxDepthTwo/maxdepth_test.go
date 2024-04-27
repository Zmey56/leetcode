package maxDepthTwo

import "testing"

// Test for function maxDepth
func Test_maxDepth(t *testing.T) {
	// Test case 1
	root1 := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	if got := maxDepth(root1); got != 3 {
		t.Errorf("maxDepth() = %v, want %v", got, 3)
	}

	// Test case 2
	root2 := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 2,
			},
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 6,
						Children: []*Node{
							{
								Val: 11,
								Children: []*Node{
									{
										Val: 14,
									},
								},
							},
						},
					},
					{
						Val: 7,
						Children: []*Node{
							{
								Val: 12,
							},
						},
					},
				},
			},
			{
				Val: 4,
				Children: []*Node{
					{
						Val: 8,
						Children: []*Node{
							{
								Val: 13,
							},
						},
					},
				},
			},
			{
				Val: 5,
				Children: []*Node{
					{
						Val: 9,
					},
					{
						Val: 10,
					},
				},
			},
		},
	}
	if got := maxDepth(root2); got != 5 {
		t.Errorf("maxDepth() = %v, want %v", got, 5)
	}

}
