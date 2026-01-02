package isHappyV2

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{19, true},
		{2, false},
	}

	for test, tt := range tests {
		result := isHappy(tt.num)
		if result != tt.want {
			t.Errorf("Test %d failed, expected %v, got %v", test, tt.want, result)
		}
	}
}

var testNumbers = []int{
	1, 7, 19, 2, 4, 20, 100, 1111111, 9999999,
}

func BenchmarkIsHappyMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isHappy(testNumbers[i%len(testNumbers)])
	}
}

func BenchmarkIsHappyFloyd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isHappyV2(testNumbers[i%len(testNumbers)])
	}
}
