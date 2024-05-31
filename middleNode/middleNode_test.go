package middleNode

import "testing"

func Test_middleNode(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "Example 1",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
		{
			name: "Example 2",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val:  6,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  6,
						Next: nil,
					},
				},
			},
		},
		{
			name: "Example 3",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
		{
			name: "Example 4",
			head: &ListNode{
				Val:  1,
				Next: nil,
			},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.head); !isEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
