package HashSet

import "testing"

// TestMyHashSet is a test function
func TestMyHashSet(t *testing.T) {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	if obj.Contains(1) != true {
		t.Errorf("Error: Expected true, got false")
	}
	if obj.Contains(3) != false {
		t.Errorf("Error: Expected false, got true")
	}
	obj.Add(2)
	if obj.Contains(2) != true {
		t.Errorf("Error: Expected true, got false")
	}
	obj.Remove(2)
	if obj.Contains(2) != false {
		t.Errorf("Error: Expected false, got true")
	}
}
