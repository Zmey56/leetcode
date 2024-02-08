package deleteDuplicates

//Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
//Return the linked list sorted as well.

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	//create slice from ListNode
	slice := createSlice(head)
	newSlice := make([]int, 0)

	var tmpValue int

	for _, value := range slice {
		if tmpValue == value {
			continue
		}
		tmpValue = value
		newSlice = append(newSlice, value)
	}

	//create ListNode from slice
	return createListNode(newSlice)
}

// function create slice from ListNode
func createSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// function create ListNode from slice
func createListNode(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}
	head := &ListNode{Val: slice[0]}
	current := head
	for i := 1; i < len(slice); i++ {
		current.Next = &ListNode{Val: slice[i]}
		current = current.Next
	}
	return head
}
