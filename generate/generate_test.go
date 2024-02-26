package generate

import "testing"

//Test for function generate

func TestGenerate(t *testing.T) {
	tests := []struct {
		name     string
		numRows  int
		expected [][]int
	}{
		{"Testcase 1", 5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{"Testcase 2", 1, [][]int{{1}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := generate(test.numRows)
			if !compare(got, test.expected) {
				t.Errorf("generate(%v) = %v, want %v", test.numRows, got, test.expected)
			}
		})
	}

}

func compare(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
