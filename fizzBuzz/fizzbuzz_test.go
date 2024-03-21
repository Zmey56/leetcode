package fizzBuzz

import (
	"reflect"
	"testing"
)

// Test for function fizzBuzz
func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{
			n:    15,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			n:    1,
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		if got := fizzBuzz(tt.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
		}
	}
}
