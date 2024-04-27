package postorder

import (
	"reflect"
	"testing"
)

// Test for function postorder

func Test_postorder(t *testing.T) {
	root := &Node{Val: 1}
	root.Children = append(root.Children, &Node{Val: 3})
	root.Children = append(root.Children, &Node{Val: 2})
	root.Children = append(root.Children, &Node{Val: 4})
	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 5})
	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 6})

	result := postorder(root)
	want := []int{5, 6, 3, 2, 4, 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("postorder(root) = %v; want %v", result, want)
	}
}
