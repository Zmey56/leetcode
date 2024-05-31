package HashMap

import "testing"

func Test_MyHashMap(t *testing.T) {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	if obj.Get(1) != 1 {
		t.Error("Get(1) != 1")
	}

	if obj.Get(3) != -1 {
		t.Error("Get(3) != -1")
	}

	obj.Put(2, 1)
	if obj.Get(2) != 1 {
		t.Error("Get(2) != 1")
	}

}
