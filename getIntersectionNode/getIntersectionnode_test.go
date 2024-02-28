package getIntersectionNode

import "testing"

//Test for function getIntersectionNode

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		headA []int
		headB []int
		want  []int
	}{
		{
			headA: []int{4, 1, 8, 4, 5},
			headB: []int{5, 0, 1, 8, 4, 5},
			want:  []int{8, 4, 5},
		},
		{
			headA: []int{0, 9, 1, 2, 4},
			headB: []int{3, 2, 4},
			want:  []int{2, 4},
		},
		{
			headA: []int{2, 6, 4},
			headB: []int{1, 5},
			want:  []int{},
		},
	}

	for _, test := range tests {
		headA := getListNode(test.headA)
		headB := getListNode(test.headB)
		got := getIntersectionNode(headA, headB)
		gotSlice := getSlice(got)
		if !compareList(gotSlice, test.want) {
			t.Errorf("getIntersectionNode(%v, %v) = %v; want %v\n", test.headA, test.headB, got, test.want)
		}

	}
}

// get the linked list from the slice
func getListNode(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode{Val: a[0]}
	p := head
	for i := 1; i < len(a); i++ {
		p.Next = &ListNode{Val: a[i]}
		p = p.Next
	}
	return head
}

// get the slice from the linked list
func getSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result

}

// compare two slice
func compareList(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
