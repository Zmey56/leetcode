package singleNumber

import "testing"

//Test for singleNumber function

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		example   string
		input     []int
		output    int
		exception bool
	}{
		{
			example:   "Test 1",
			input:     []int{2, 2, 1},
			output:    1,
			exception: true,
		},
		{
			example:   "Test 2",
			input:     []int{4, 1, 2, 1, 2},
			output:    4,
			exception: true,
		},
		{
			example:   "Test 3",
			input:     []int{1},
			output:    1,
			exception: true,
		},
	}
	for _, test := range tests {
		result := singleNumber(test.input)
		if result != test.output {
			t.Errorf("Test %s failed, expected %v, got %v", test.example, test.output, result)
		}
	}
}
