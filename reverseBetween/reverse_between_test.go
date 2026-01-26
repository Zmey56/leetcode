package reversebetween

import (
	"reflect"
	"testing"
)

// sliceToList — помощник для создания списка из слайса
func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

// listToSlice — помощник для проверки результата
func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		left  int
		right int
		want  []int
	}{
		{"Example 1", []int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{"Example 2", []int{5}, 1, 1, []int{5}},
		{"Full reverse", []int{1, 2, 3}, 1, 3, []int{3, 2, 1}},
		{"Reverse first two", []int{1, 2, 3}, 1, 2, []int{2, 1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.input)
			gotHead := reverseBetween(head, tt.left, tt.right)
			got := listToSlice(gotHead)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReverseBetween(b *testing.B) {
	// Создаем список из 500 элементов (максимум по условию)
	nums := make([]int, 500)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		head := sliceToList(nums)
		reverseBetween(head, 1, 500)
	}
}
