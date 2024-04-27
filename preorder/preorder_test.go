package preorder

import (
	"reflect"
	"testing"
)

// Test for function preorder
func Test_preorder(t *testing.T) {
	root := &Node{Val: 1}
	root.Children = append(root.Children, &Node{Val: 3})
	root.Children = append(root.Children, &Node{Val: 2})
	root.Children = append(root.Children, &Node{Val: 4})
	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 5})
	root.Children[0].Children = append(root.Children[0].Children, &Node{Val: 6})

	result := preorder(root)
	want := []int{1, 3, 5, 6, 2, 4}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("preorder(root) = %v; want %v", result, want)
	}
}

//
