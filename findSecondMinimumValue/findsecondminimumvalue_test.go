package findSecondMinimumValue

import "testing"

func Test_findSecondMinimumValue(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			want: 5,
		},
		{
			name: "Test case 2",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: -1,
		},
		{
			name: "Test case 3",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSecondMinimumValue(tt.args.root); got != tt.want {
				t.Errorf("findSecondMinimumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
