package calculate

import (
	"testing" // Обязательно добавьте этот импорт
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1 + 1", 2},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
	}

	for _, test := range tests {
		result := calculate(test.input)
		if result != test.expected {
			t.Errorf("calculate(%q) = %d, want %d", test.input, result, test.expected)
		}
	}
}
