package isPalindrome

import "testing"

//Test for isPalindrome function

func TestIsPalindrom(t *testing.T) {
	tests := []struct {
		example   string
		input     string
		output    bool
		exception bool
	}{
		{
			example:   "Test 1",
			input:     "A man, a plan, a canal: Panama",
			output:    true,
			exception: true,
		},
		{
			example:   "Test 2",
			input:     "race a car",
			output:    false,
			exception: false,
		},
		{
			example:   "Test 3",
			input:     " ",
			output:    true,
			exception: true,
		},
		{
			example:   "Test 4",
			input:     "0P",
			output:    false,
			exception: false,
		},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.output {
			t.Errorf("Test %s failed, expected %v, got %v", test.example, test.output, result)
		}
	}

}
