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
	dividend6 := 1
	divisor6 := -1
	expected6 := -1
	result6 := divide(dividend6, divisor6)
	if result6 != expected6 {
		t.Errorf("Expected %d, but got %d", expected6, result6)
	}

}
