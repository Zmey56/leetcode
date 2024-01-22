package divide

import "testing"

//Create test for function divide

func TestDivide(t *testing.T) {
	//Example 1
	dividend1 := 10
	divisor1 := 3
	expected1 := 3
	result1 := divide(dividend1, divisor1)
	if result1 != expected1 {
		t.Errorf("Expected %d, but got %d", expected1, result1)
	}

	//Example 2
	dividend2 := 7
	divisor2 := -3
	expected2 := -2
	result2 := divide(dividend2, divisor2)
	if result2 != expected2 {
		t.Errorf("Expected %d, but got %d", expected2, result2)
	}

	//Example 3
	dividend3 := 0
	divisor3 := 1
	expected3 := 0
	result3 := divide(dividend3, divisor3)
	if result3 != expected3 {
		t.Errorf("Expected %d, but got %d", expected3, result3)
	}

	//Example 4
	dividend4 := 1
	divisor4 := 1
	expected4 := 1
	result4 := divide(dividend4, divisor4)
	if result4 != expected4 {
		t.Errorf("Expected %d, but got %d", expected4, result4)
	}

	//Example 5
	dividend5 := -2147483648
	divisor5 := -1
	expected5 := 2147483647
	result5 := divide(dividend5, divisor5)
	if result5 != expected5 {
		t.Errorf("Expected %d, but got %d", expected5, result5)
	}

	//Example 6
	dividend6 := -2147483648
	divisor6 := 1
	expected6 := -2147483648
	result6 := divide(dividend6, divisor6)
	if result6 != expected6 {
		t.Errorf("Expected %d, but got %d", expected6, result6)
	}

	//Example 7
	dividend7 := -2147483648
	divisor7 := -2147483648
	expected7 := 1
	result7 := divide(dividend7, divisor7)
	if result7 != expected7 {
		t.Errorf("Expected %d, but got %d", expected7, result7)
	}

	//Example 8
	dividend8 := 10
	divisor8 := 3
	expected8 := 3
	result8 := divide(dividend8, divisor8)
	if result8 != expected8 {
		t.Errorf("Expected %d, but got %d", expected8, result8)
	}

	//Example 9
	dividend9 := 7
	divisor9 := -3
	expected9 := -2
	result9 := divide(dividend9, divisor9)
	if result9 != expected9 {
		t.Errorf("Expected %d, but got %d", expected9, result9)
	}
}
