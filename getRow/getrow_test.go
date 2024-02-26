package getRow

import "testing"

//Test for function getRow

func TestGetRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		expected []int
	}{
		{"Testcase 1", 3, []int{1, 3, 3, 1}},
		{"Testcase 2", 0, []int{1}},
		{"Testcase 3", 1, []int{1, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getRow(test.rowIndex)
			if !compare(got, test.expected) {
				t.Errorf("getRow(%v) = %v, want %v", test.rowIndex, got, test.expected)
			}
		})
	}

}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
