package reverseString

import (
	"reflect"
	"testing"
)

//Test for reverseString

func TestReverseString(t *testing.T) {

	test := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		{[]byte{'a', 'b', 'c', 'd', 'e', 'f'}, []byte{'f', 'e', 'd', 'c', 'b', 'a'}},
	}

	for _, test := range test {
		reverseString(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("Test failed: expected %v, got %v", test.expected, test.input)
		}
	}
}
